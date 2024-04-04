#include "TextureManager.h"
#include "Game.h"
#include "SDL2/SDL_ttf.h"
#include <iostream>

SDL_Texture *TextureManager::LoadTexture(const char *fileName) {
    SDL_Surface* tmpSurface = IMG_Load(fileName);
    SDL_Texture * tex = SDL_CreateTextureFromSurface(Game::renderer, tmpSurface);
    SDL_FreeSurface(tmpSurface);

    return tex;
}

void TextureManager::Draw(SDL_Texture *tex, SDL_Rect src, SDL_Rect dst) {
    SDL_RenderCopy(Game::renderer, tex, &src, &dst);
}

// Creates texture with text using ttf font
// Parameters:
//    path_to_ttf: path to ttf file
//    size: size of text
//    color: color of text
//    text: message that will be in texture
//    rect: pointer to rect (w and h will be set to surface w and h)
// Returns:
//    texture with text
SDL_Texture *TextureManager::LoadTTFTexture(const char *path_to_ttf, int size, SDL_Color color, const char* text, SDL_Rect& rect) {
    // Load font
    TTF_Font * font = TTF_OpenFont(path_to_ttf, size);
    if (!font) {
        std::cerr << "Failed to load font: " << TTF_GetError() << std::endl;
        return nullptr;
    }

    // Render text to surface
    SDL_Surface * tmpSurface = TTF_RenderText_Solid(font, text, color);
    if (!tmpSurface) {
        std::cerr << "Failed to render text: " << TTF_GetError() << std::endl;
        TTF_CloseFont(font);
        return nullptr;
    }

    // Create texture from surface
    SDL_Texture * tex = SDL_CreateTextureFromSurface(Game::renderer, tmpSurface);
    if (!tex) {
        std::cerr << "Failed to create texture from surface: " << SDL_GetError() << std::endl;
        SDL_FreeSurface(tmpSurface);
        TTF_CloseFont(font);
        return nullptr;
    }

    // Set rectangle dimensions
    rect.h = tmpSurface->h;
    rect.w = tmpSurface->w;

    // Clean up resources
    SDL_FreeSurface(tmpSurface);
    TTF_CloseFont(font);

    return tex;
}
