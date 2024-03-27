//
// Created by a on 27.03.2024.
//

#ifndef GAME_SDL2_TYPES_H
#define GAME_SDL2_TYPES_H

#include "Math.hpp"

struct Player {
    char username[20]{};
    uint16_t id;
    Vector2f pos;

    Player() :username(""), id(0), pos(Vector2f())  {};

//    uint32_t pin;
};

#endif //GAME_SDL2_TYPES_H
