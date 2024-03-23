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
auto& newPlayer(manager.addEntity());

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

    newPlayer.addComponent<TransformComponent>();
    newPlayer.addComponent<SpriteComponent>("res/imgs/star.png");
}

void Game::update() {
    manager.update();

    std::cout << "Player pos: " << newPlayer.getComponent<TransformComponent>().pos() << std::endl;
}

void Game::handleEvents() {
    SDL_Event event;
    SDL_PollEvent(&event);
    switch (event.type) {
        case SDL_QUIT:
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
    SDL_DestroyWindow(window);
    SDL_DestroyRenderer(renderer);
    SDL_Quit();
    std::cout << "Game cleaned" << std::endl;
}