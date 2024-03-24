//
// Created by a on 22.03.2024.
//

#ifndef GAME_SDL2_GAME_H
#define GAME_SDL2_GAME_H

#include <SDL.h>
#include <SDL_image.h>
#include <vector>

class BoxColliderComponent;

class Game {
public:
    Game();
    ~Game();

    void init(const char * title, int xpos, int ypos, int width, int height, bool fullscreen);

    void handleEvents();
    static void update();
    static void render();
    void clean();

    [[nodiscard]] bool running() const {return isRunning;}

    static SDL_Renderer *renderer;
    static SDL_Event event;

    static std::vector<BoxColliderComponent*> colliders;
private:
    bool isRunning{};
    SDL_Window *window{};
};


#endif //GAME_SDL2_GAME_H
