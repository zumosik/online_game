//
// Created by a on 22.03.2024.
//

#ifndef GAME_SDL2_TEXTUREMANAGER_H
#define GAME_SDL2_TEXTUREMANAGER_H

#include <SDL.h>
#include "SDL_image.h"

class TextureManager {
public:
    static SDL_Texture* LoadTexture(const char *fileName);
    static void Draw(SDL_Texture* tex, SDL_Rect src,SDL_Rect dst);
};

#endif //GAME_SDL2_TEXTUREMANAGER_H
