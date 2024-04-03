//
// Created by a on 24.03.2024.
//

#include "Collision.h"

bool Collision::AABB(const BoxColliderComponent *colA, const BoxColliderComponent *colB) {
    if (colA && colB && // if this is null game just crashes
        colA->tag != colB->tag && // not sure about this
        SDL_HasIntersection(&colA->collider, &colB->collider))
    {
//        std::cout << colA->tag << " hit " << colB->tag << std::endl;
        return true;
    }

    return false;
}