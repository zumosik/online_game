//
// Created by a on 22.03.2024.
//

#ifndef GAME_SDL2_GAME_H
#define GAME_SDL2_GAME_H

#include <SDL.h>
#include <SDL_image.h>


class Game {
public:
    Game();
    ~Game();

    void init(const char * title, int xpos, int ypos, int width, int height, bool fullscreen);

    void handleEvents();
    void update();
    void render();
    void clean();

    bool running() {return isRunning;}
private:
    int cnt = 0;
    bool isRunning{};
    SDL_Window *window{};
    SDL_Renderer *renderer{};
};


#endif //GAME_SDL2_GAME_H
