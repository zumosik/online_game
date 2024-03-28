//
// Created by a on 23.03.2024.
//

#ifndef GAME_SDL2_TRANSFORMCOMPONENT_H
#define GAME_SDL2_TRANSFORMCOMPONENT_H

#include "ECS/Components.h"
#include "ECS/ECS.h"
#include "Types.hpp"

class TransformComponent : public Component {
public:
    Vector2f position;
    Vector2f velocity;
    int h = 32;
    int w = 32;
    int scale = 1;

    float speed = 0;

    TransformComponent() {position.Zero(); }
    explicit TransformComponent(Vector2f p_pos, float p_speed) : position(p_pos), speed(p_speed){}
    explicit TransformComponent(Vector2f p_pos, float p_speed, int p_h, int p_w) : position(p_pos), speed(p_speed), h(p_h), w(p_w){}
    explicit TransformComponent(Vector2f p_pos, float p_speed, int p_h, int p_w, int p_scale) : position(p_pos), speed(p_speed), h(p_h), w(p_w), scale(p_scale){}


    void init() override {
        velocity.Zero();
    }

    void update() override {
        position.x += velocity.x * speed;
        position.y += velocity.y * speed;
    }
};

#endif //GAME_SDL2_TRANSFORMCOMPONENT_H
