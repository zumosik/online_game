//
// Created by a on 27.03.2024.
//

#ifndef GAME_SDL2_PLAYERINFOCOMPONENT_H
#define GAME_SDL2_PLAYERINFOCOMPONENT_H

#include "ECS/Components.h"

class PlayerInfoComponent : public Component {
public:
    std::string username;
    uint16_t id;
    Vector2f startPos;

    PlayerInfoComponent( char p_username[20],  uint16_t p_id, Vector2f p_std_pos) {
        username = p_username;
        id = p_id;
        startPos = p_std_pos;

    }

    void init() override {
        if (!(entity->hasComponent<TransformComponent>())) entity->addComponent<TransformComponent>();
        TransformComponent* transform = &entity ->getComponent<TransformComponent>();
        transform->position = startPos; // setting start pos
    }
};

#endif //GAME_SDL2_PLAYERINFOCOMPONENT_H
