//
// Created by a on 24.03.2024.
//

#ifndef GAME_SDL2_COLLISION_H
#define GAME_SDL2_COLLISION_H

#include "ECS/Components.h"

class Collision {
public:
    static bool AABB(const BoxColliderComponent* colA, const BoxColliderComponent* colB);
};


#endif //GAME_SDL2_COLLISION_H
