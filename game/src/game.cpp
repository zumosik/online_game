#include <vector>

#include "Game.h"
#include "Utils.hpp"

void Game::GameLoop() {
    SDL_Texture* grassTexture = window.loadTexture("res/imgs/ground_grass_1.png");


    for (int i = 0; i < 30; ++i)
        entities.emplace_back(Vector2f(100 + (i * 32), 100), grassTexture);


    SDL_Event event;

    const float deltaTime = 1.0f/120.0f; // 120 fps
    float accumulator = 0.0f;
    float currTime = utils::hireTimeInSeconds();

    int refreshRate = window.getRefreshRate();

    std::cout << "Starting game..." << std::endl;
    std::cout << "FPS: " << refreshRate << std::endl;

    while (gameRunning) {
        Uint32 startTicks = SDL_GetTicks();

        float newTime = utils::hireTimeInSeconds();
        float frameTime = newTime - currTime;

        currTime = newTime;

        accumulator += frameTime;

        handleEvents();


        while (accumulator >= deltaTime) {
            while (SDL_PollEvent(&event)) {
                if (event.type == SDL_QUIT)
                    gameRunning = false;
            }

            accumulator -= deltaTime;
        }
//        const float alpha = accumulator / deltaTime;
        handleFrame();
        Uint32 frameTicks = SDL_GetTicks() - startTicks;

        if (frameTicks < (1000/ refreshRate)) SDL_Delay((1000 / refreshRate) - frameTicks);
    }
}

void Game::handleFrame() {
    window.clear();


    player.Move(); // apply velocity


    for (Entity & entity : entities) {
        window.render(entity);
    }

    window.render(player);

    window.display();
}

void Game::handleEvents() {
    SDL_Event event;
    while (SDL_PollEvent(&event)) {
        switch (event.type) {
            case SDL_QUIT:
                gameRunning = false;
                break;
            case SDL_KEYDOWN:
                handleKeyDown(event); // Pass the event object
                break;
            case SDL_KEYUP:
                handleKeyUp(event); // Pass the event object
                break;
            default:
                break;
        }
    }
}

void Game::handleKeyDown(SDL_Event event) {
    switch (event.key.keysym.sym) {
        case SDLK_w:
            player.velocity.y = -1.0f;
            break;
        case SDLK_s:
            player.velocity.y = 1.0f;
            break;
        case SDLK_a:
            player.velocity.x = -1.0f;
            break;
        case SDLK_d:
            player.velocity.x = 1.0f;
            break;
        default:
            break;
    }
}

void Game::handleKeyUp(SDL_Event event) {
    switch (event.key.keysym.sym) {
        case SDLK_w:

            player.velocity.y = 0.0f;
            break;
        case SDLK_s:
            player.velocity.y = 0.0f;
            break;
        case SDLK_a:
            player.velocity.x = 0.0f;
            break;
        case SDLK_d:
            player.velocity.x = 0.0f;
            break;
        default:
            break;
    }
}
