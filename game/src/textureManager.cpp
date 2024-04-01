#include "TextureManager.h"
#include "Game.h"
#include "SDL_ttf.h"

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
//    w pointer: will be changed to surface w
//    h pointer: will be changed to surface h
// Returns:
//    texture with text
SDL_Texture *TextureManager::LoadTTFTexture(const char *path_to_ttf, int size, SDL_Color color, const char* text, int& w, int& h) {
    TTF_Font * font = TTF_OpenFont(path_to_ttf, size);
    SDL_Surface * tmpSurface = TTF_RenderText_Solid(font, text, color);
    SDL_Texture * tex = SDL_CreateTextureFromSurface(Game::renderer, tmpSurface);

    w = tmpSurface->w;
    h = tmpSurface->h;

    SDL_FreeSurface(tmpSurface);
    return tex;
}
