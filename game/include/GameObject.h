//
// Created by a on 23.03.2024.
//

#ifndef GAME_SDL2_GAMEOBJECT_H
#define GAME_SDL2_GAMEOBJECT_H

#include "SDL.h"
#include "Math.hpp"

class GameObject {
public:
    GameObject(const char* textureSheet, SDL_Renderer *ren);
    ~GameObject();

    void Update();
    void Render();
private:
    Vector2f pos;

    SDL_Texture* texture;
    SDL_Rect srcRect, dstRect;
    SDL_Renderer * renderer;
};


#endif //GAME_SDL2_GAMEOBJECT_H
