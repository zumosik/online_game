#include <SDL.h>
#include <SDL_image.h>
#include <iostream>
#include <winsock2.h> // Very important to compile code
#include <windows.h> // Windows support
#include <thread>

#include "RenderWindow.hpp"
#include "TCPClient.hpp"
#include "Game.h"

boost::asio::io_context io_context;
Game *game = nullptr;

void runIoContext() {
    io_context.run();
}

// Define WinMain as the entry point for Windows GUI applications
int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int nCmdShow) {
    // Init sdl...
//    if (SDL_Init(SDL_INIT_VIDEO) > 0)
//        std::cout << "SDL_Init has failed. SDL_ERROR: " << SDL_GetError() << std::endl;
//
//    if (!IMG_Init(IMG_INIT_PNG))
//        std::cout << "IMG_Init has failed. Error: " << SDL_GetError() << std::endl;

    // TODO uncomment
/*
    // Connect to TCP
    auto addr = "127.0.0.1";
    int port = 8080;
    std::cout << "Connecting to: " << addr << ":" << port << std::endl;
    TCPClient tcpClient(io_context);
    tcpClient.start(addr, port, io_context);

    std::thread ioContextThread(runIoContext); // Run io_context.run() in a separate thread

    */

    // All Game stuff
    const int FPS = 60;
    const int frameDelay = 1000/FPS;
    std::cout << "Running game in: " << FPS << " fps" << std::endl;

    Uint32 frameStart;
    int frameTime;

    game = new Game();

    game->init("Game", SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, 1280,720, false);

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

    // TODO uncomment
/*
    tcpClient.stop();
    io_context.stop(); // Stop io_context when program exits
    ioContextThread.join(); // Wait for io_context thread to finish

    */

    return 0;
}
