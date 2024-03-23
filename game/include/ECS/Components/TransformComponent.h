//
// Created by a on 23.03.2024.
//

#ifndef GAME_SDL2_TRANSFORMCOMPONENT_H
#define GAME_SDL2_TRANSFORMCOMPONENT_H

#include "ECS/Components.h"
#include "ECS/ECS.h"
#include "Math.hpp"

class TransformComponent : public Component {
public:
    Vector2f position;
    Vector2f velocity;

    float speed = 3;

    TransformComponent() {position = Vector2f(); }
    explicit TransformComponent(Vector2f pos, float speed) : position(pos), speed(speed){}


    void init() override {
        velocity = Vector2f();
    }

    void update() override {
        position.x += velocity.x * speed;
        position.y += velocity.y * speed;
    }
};

#endif //GAME_SDL2_TRANSFORMCOMPONENT_H
