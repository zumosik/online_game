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
    TransformComponent() {position = Vector2f();}
    explicit TransformComponent(Vector2f pos) : position(pos){}

    Vector2f pos() { return position;}
    void setPos(Vector2f pos){position = pos;}

    void init() override {
        position = Vector2f();
    }

    void update() override {
        position + Vector2f(1,1);
    }
private:
    Vector2f position;
};

#endif //GAME_SDL2_TRANSFORMCOMPONENT_H
