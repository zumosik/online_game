//
// Created by a on 22.03.2024.
//

#include <iostream>

#include "TextureManager.h"
#include "Game.h"
#include "Map.h"

#include "ECS/ECS.h"
#include "ECS/Components.h"
#include "Collision.h"

#include "SDL_ttf.h"

enum groupLabels : std::size_t {
    GROUP_MAP,
    GROUP_PLAYERS,
    GROUP_UI,
};

//Map *map;

SDL_Renderer* Game::renderer = nullptr;
std::chrono::milliseconds Game::ping = std::chrono::milliseconds();
std::chrono::milliseconds Game::prev_ping = std::chrono::milliseconds();

Manager manager;
SDL_Event Game::event;

std::vector<BoxColliderComponent*> Game::colliders;

Entity& player(manager.addEntity());

std::vector<Entity*> otherPlayers;

Entity& wall(manager.addEntity());

Entity& text (manager.addEntity());

Entity& tile10(manager.addEntity());
Entity& tile11(manager.addEntity());
Entity& tile12(manager.addEntity());


Game::Game() = default;
Game::~Game() = default;

void Game::init(const char *title, int xpos, int ypos, int width, int height, bool fullscreen) {
    int flags = 0;

    if (fullscreen) flags = SDL_WINDOW_FULLSCREEN;

    if (SDL_Init(SDL_INIT_EVERYTHING) == 0) {
        std::cout << "Sdl initiated" << std::endl;
    } else {
        std::cerr << "Can initialize SDL: " << SDL_GetError() << std::endl;
        isRunning = false;
        return;
    }


    if (TTF_Init() == 0) {
        std::cout << "Sdl ttf initiated" << std::endl;
    } else {
        std::cerr << "Can initialize SDL_TTF: " << SDL_GetError() << std::endl;
        isRunning = false;
        return;
    }



    if (!IMG_Init(IMG_INIT_PNG))
        std::cout << "IMG_Init has failed. Error: " << SDL_GetError() << std::endl;

    window = SDL_CreateWindow(title, xpos, ypos, width, height, flags);
    if (window) {
        std::cout << "Window created" << std::endl;
    }

    renderer = SDL_CreateRenderer(window, -1, 0);
    if (renderer) {
        SDL_SetRenderDrawColor(renderer, 255,255,255,255);
        std::cout << "Renderer created" << std::endl;
    }

    isRunning = true;
//    map = new Map();

    tile10.addComponent<TileComponent>(200,200,32,32,GRASS);
    tile10.addGroup(GROUP_MAP);

    tile11.addComponent<TileComponent>(250,250,32,32,GRASS);
    tile11.addGroup(GROUP_MAP);

    tile12.addComponent<TileComponent>(300,300,32,32,GRASS);
    tile12.addGroup(GROUP_MAP);

    tile10.addComponent<BoxColliderComponent>("grass");
    tile11.addComponent<BoxColliderComponent>("grass");
    tile12.addComponent<BoxColliderComponent>("grass");

    player.addComponent<TransformComponent>(Vector2f(), 5.0f, 32, 32, 2);
    player.addComponent<SpriteComponent>("res/imgs/star.png");
    player.addComponent<KeyboardControllerComponent>();
    player.addComponent<BoxColliderComponent>("player");
    player.addGroup(GROUP_PLAYERS);


    SDL_Color col = {0,0,0};
    text.addComponent<TransformComponent>(Vector2f(0,400), 0, 32, 32, 2);
    if (!text.hasComponent<SpriteTextComponent>()) {
        text.addComponent<SpriteTextComponent>("res/font.ttf", 32, col, "");
    }
    text.getComponent<SpriteTextComponent>().setText("text");
    text.addGroup(GROUP_UI);

//
//    wall.addComponent<TransformComponent>(Vector2f(300,300),0, 300, 20, 1);
//    wall.addComponent<SpriteComponent>("res/imgs/grass.png");
//    wall.addComponent<BoxColliderComponent>("wall");

}

void Game::update() {
    manager.refresh();
    manager.update();

    if (ping != prev_ping) {
        std::string str =  "Ping: " + std::to_string(ping.count()) + "ms";
        text.getComponent<SpriteTextComponent>().setText(str.c_str());
        prev_ping = ping;
    }


    for (auto c : colliders) {
        if (Collision::AABB(&player.getComponent<BoxColliderComponent>(), c))
            player.getComponent<TransformComponent>().velocity * -1; // pushing player back
    }

//    std::cout << "Player pos: " << player.getComponent<TransformComponent>().position << std::endl;
}

void Game::handleEvents() {
    SDL_PollEvent(&event);
    switch (event.type) {
        case SDL_QUIT:
            std::cout << "SDL QUIT" << std::endl;
            isRunning = false;
            break;
        default:
            break;
    }
}


auto& tiles(manager.getGroup(GROUP_MAP));
auto& players(manager.getGroup(GROUP_PLAYERS));
auto& ui(manager.getGroup(GROUP_UI));

void Game::render() {
    SDL_RenderClear(renderer);

//    map->DrawMap();
//    manager.draw();

    for (auto& t: tiles) {
        t->draw();
    }

    for(auto& p : players) {
        p->draw();
    }

    for (auto & e : ui) {
        e->draw();
    }

    SDL_RenderPresent(renderer);

}
void Game::clean() {
    std::cout << "Cleaning.." << std::endl;
    SDL_DestroyWindow(window);
    SDL_DestroyRenderer(renderer);
    SDL_Quit();
    std::cout << "Game cleaned" << std::endl;
}

void Game::InitializePlayer(ConnectResp *resp) {
    player.addComponent<PlayerInfoComponent>(resp->player.username, resp->player.id, resp->player.pos);
    std::cout << "InitializePlayer, id = " << resp->player.id   << std::endl;
}

void Game::SetPing(std::chrono::milliseconds duration) {
    ping = duration;
}

void Game::SpawnNewPlayer(NewPlayerConnect *req) {
//    otherPlayers.push_back(&player);
    Entity& otherPlayer(manager.addEntity());
    player.addComponent<TransformComponent>(req->player.pos, 0, 32, 32);
    player.addComponent<SpriteComponent>("res/imgs/star.png");
    player.addGroup(GROUP_PLAYERS);

    otherPlayers.push_back(&otherPlayer);
}
