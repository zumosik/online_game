//
// Created by a on 22.03.2024.
//

#ifndef GAME_SDL2_GAME_H
#define GAME_SDL2_GAME_H

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <vector>
#include "Packet.h"

#include <chrono>



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

    static void InitializePlayer(ConnectResp* req);

    static void SpawnNewPlayer(NewPlayerConnect* req);

    static void SetPing(std::chrono::milliseconds duration);

private:
    static std::chrono::milliseconds  ping;
    static std::chrono::milliseconds  prev_ping;
    bool isRunning{};
    SDL_Window *window{};
};


#endif //GAME_SDL2_GAME_H
