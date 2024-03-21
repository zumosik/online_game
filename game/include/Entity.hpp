#pragma once
#include <SDL.h>
#include <SDL_image.h>

#include "Math.hpp"

class Entity
{
public:
    Entity(Vector2f p_pos, SDL_Texture* p_tex);
    Vector2f& getPos();
    void setPos(Vector2f& newPos);
    SDL_Rect getCurrentFrame();
    SDL_Texture* getTex();
private:
    Vector2f pos;
    SDL_Rect currentFrame;
    SDL_Texture* tex;
};