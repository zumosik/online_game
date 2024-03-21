#pragma once

#include "RenderWindow.hpp"
#include "Player.hpp"


class Game {
public:
    explicit Game(RenderWindow & window, Player player) : window(window), player(player), entities() {}

    void GameLoop();
private:
    void handleEvents();
    void handleKeyDown(SDL_Event event);
    void handleKeyUp(SDL_Event event);
    void handleFrame();

    RenderWindow window;
    bool gameRunning = true;
    Player player;
    std::vector<Entity> entities;
};


