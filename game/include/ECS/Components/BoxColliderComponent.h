//
// Created by a on 24.03.2024.
//

#ifndef GAME_SDL2_BOXCOLLIDERCOMPONENT_H
#define GAME_SDL2_BOXCOLLIDERCOMPONENT_H

#include <string>
#include "SDL.h"
#include "ECS/Components.h"

class BoxColliderComponent : public Component{
public:
    SDL_Rect collider;
    std::string tag;
    TransformComponent* transform;

    explicit BoxColliderComponent(std::string p_tag) : tag(std::move(p_tag)){}

    void init() override {
        if (!(entity->hasComponent<TransformComponent>())) entity->addComponent<TransformComponent>();
        transform = &entity ->getComponent<TransformComponent>();
    }

    void update() override {
        collider.x = static_cast<int>(transform->position.x);
        collider.y = static_cast<int>(transform->position.y);
        collider.w = transform->w * transform->scale;
        collider.h = transform->h * transform->scale;
    }


};

#endif //GAME_SDL2_BOXCOLLIDERCOMPONENT_H
