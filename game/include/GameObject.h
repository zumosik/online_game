//
// Created by a on 23.03.2024.
//

#ifndef GAME_SDL2_GAMEOBJECT_H
#define GAME_SDL2_GAMEOBJECT_H

#include "SDL.h"
#include "Math.hpp"

class GameObject {
public:
    GameObject(const char* textureSheet,  Vector2int pos);
    ~GameObject();

    void Update();
    void Render();
private:
    Vector2int pos;

    SDL_Texture* texture;
    SDL_Rect srcRect{}, dstRect{};
};


#endif //GAME_SDL2_GAMEOBJECT_H
