//
// Created by a on 23.03.2024.
//

#ifndef GAME_SDL2_MAP_H
#define GAME_SDL2_MAP_H


#include <SDL.h>

class Map {
public:
    Map();
    ~Map();

    void LoadMap(int arr[20][25]);
    void DrawMap();

private:
    SDL_Rect src,dst;
    SDL_Texture *grass;

    uint8_t map[20][25]; // TODO change map size
};


#endif //GAME_SDL2_MAP_H
