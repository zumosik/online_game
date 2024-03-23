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
    TransformComponent() {position = Vector2int();}
    explicit TransformComponent(Vector2int pos) : position(pos){}

    Vector2int pos() { return position;}
    void setPos(Vector2int pos){position = pos;}

    void init() override {
        position = Vector2int();
    }

    void update() override {
        position.add(Vector2int(1,1));
    }
private:
    Vector2int position;
};

#endif //GAME_SDL2_TRANSFORMCOMPONENT_H
