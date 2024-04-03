//
// Created by a on 23.03.2024.
//

#ifndef GAME_SDL2_KEYBOARDCONTROLLERCOMPONENT_H
#define GAME_SDL2_KEYBOARDCONTROLLERCOMPONENT_H

#include "ECS/ECS.h"
#include "TransformComponent.h"
#include "Game.h"

class KeyboardControllerComponent  : public Component {
public:
    TransformComponent *transform;

    void init() override {
        transform  = &entity->getComponent<TransformComponent>();
    }

    void update() override {
        if (Game::event.type == SDL_KEYDOWN) {
            switch (Game::event.key.keysym.sym) {
                case SDLK_w:
                    transform->velocity.y = -1.0f;
                    break;
                case SDLK_s:
                    transform->velocity.y = 1.0f;
                    break;
                case SDLK_a:
                    transform->velocity.x = -1.0f;
                    break;
                case SDLK_d:
                    transform->velocity.x = 1.0f;
                    break;
            }
        }

        if (Game::event.type == SDL_KEYUP) {
            switch (Game::event.key.keysym.sym) {
                case SDLK_w:
                    transform->velocity.y = 0.0f;
                    break;
                case SDLK_s:
                    transform->velocity.y = 0.0f;
                    break;
                case SDLK_a:
                    transform->velocity.x = 0.0f;
                    break;
                case SDLK_d:
                    transform->velocity.x = 0.0f;
                    break;
            }
        }
    }

};

#endif //GAME_SDL2_KEYBOARDCONTROLLERCOMPONENT_H
