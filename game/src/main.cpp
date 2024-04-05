#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_ttf.h>

#include <iostream>
#include <thread>
#include "TCPClient.h"
#include "Game.h"

boost::asio::io_context io_context;
Game *game = nullptr;

void runIoContext() {
    io_context.run();
}

int main(int argc, char* argv[]) {
    // --- Get name and pin from flags ---
    
    std::string name;
    uint32_t pin;

    if (argc < 2) {
        std::cerr << "You need to pass name as first arg and pin as second" << std::endl;
        return 1;
    }

    name = argv[1];
    pin = static_cast<uint32_t>(std::stoul(argv[2]));

    if (SDL_Init(SDL_INIT_EVERYTHING) == 0) {
        std::cout << "Sdl initiated" << std::endl;
    } else {
        std::cerr << "Cant initialize SDL: " << SDL_GetError() << std::endl;
        return 1;
    }

    // --- Init SDL ---
    
    if (!IMG_Init(IMG_INIT_PNG)) {
        std::cout << "IMG_Init has failed. Error: " << SDL_GetError() << std::endl;
        return 1;
    } else {
        std::cout << "Sdl image initiated" << std::endl;
    }

    if (TTF_Init() == 0) {
        std::cout << "Sdl ttf initiated" << std::endl;
    } else {
        std::cerr << "Can initialize SDL_TTF: " << SDL_GetError() << std::endl;
        return 1;
    }


    // --- Connect to TCP ---
    std::string addr = "172.18.254.186";
    int port = 8080;
    std::cout << "Connecting to: " << addr << ":" << port << std::endl;
    TCPClient tcpClient(io_context, name, pin);
    tcpClient.start(addr, port, io_context);

    std::thread ioContextThread(runIoContext); // Run io_context.run() in a separate thread


    // --- All Game stuff ---
    const int FPS = 60;
    const int frameDelay = 1000/FPS;
    std::cout << "Starting game in " << FPS << " fps" << std::endl;

    Uint32 frameStart;
    int frameTime;

    game = new Game();

    game->init("Game", SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, 1280,720, false);

    std::cout << "starting game" << std::endl;

    while (game->running()) {
        frameStart = SDL_GetTicks();

        game->handleEvents();
        game->update();
        game->render();

        frameTime  = SDL_GetTicks() - frameStart;

        if (frameDelay > frameTime)
            SDL_Delay(frameDelay - frameTime);

    }
    game->clean();


    tcpClient.stop();
    io_context.stop(); // Stop io_context when program exits
    ioContextThread.join(); // Wait for io_context thread to finish



    return 0;
}
