//
// Created by a on 22.03.2024.
//

#include <iostream>

#include "TextureManager.h"
#include "Game.h"
#include "Map.h"

#include "ECS/ECS.h"
#include "ECS/Components.h"


Map *map;

SDL_Renderer* Game::renderer = nullptr;

Manager manager;
SDL_Event Game::event;

auto& player(manager.addEntity());
auto& wall(manager.addEntity());

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
    map = new Map();

    player.addComponent<TransformComponent>(Vector2f(), 5.0f, 32, 32, 2);
    player.addComponent<SpriteComponent>("res/imgs/star.png");
    player.addComponent<KeyboardControllerComponent>();
    player.addComponent<BoxColliderComponent>("player");

    wall.addComponent<TransformComponent>(Vector2f(300,300),0, 300, 20, 1);
    wall.addComponent<SpriteComponent>("res/imgs/grass.png");
    wall.addComponent<BoxColliderComponent>("wall");

}

void Game::update() {
    manager.update();


    BoxColliderComponent* playerCollider = &player.getComponent<BoxColliderComponent>();
    BoxColliderComponent* wallCollider = &wall.getComponent<BoxColliderComponent>();

    if (playerCollider && wallCollider && // if this is null game just crashes
        SDL_HasIntersection(&playerCollider->collider, &wallCollider->collider))
    {
        std::cout << "WALL!" << std::endl;
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
void Game::render() {
    SDL_RenderClear(renderer);

    map->DrawMap();
    manager.draw();

    SDL_RenderPresent(renderer);

}
void Game::clean() {
    std::cout << "Cleaning.." << std::endl;
    SDL_DestroyWindow(window);
    SDL_DestroyRenderer(renderer);
    SDL_Quit();
    std::cout << "Game cleaned" << std::endl;
}