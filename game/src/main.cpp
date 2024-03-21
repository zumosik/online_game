#include <SDL.h>
#include <SDL_image.h>
#include <iostream>
#include <winsock2.h> // Very important to compile code
#include <windows.h> // Windows support
#include <thread>

#include "RenderWindow.hpp"
#include "TCPClient.hpp"
#include "Game.h"

boost::asio::io_context io_context; // Move io_context declaration outside WinMain

void runIoContext() {
    io_context.run();
}

// Define WinMain as the entry point for Windows GUI applications
int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int nCmdShow) {
    // Init sdl...
    if (SDL_Init(SDL_INIT_VIDEO) > 0)
        std::cout << "SDL_Init has failed. SDL_ERROR: " << SDL_GetError() << std::endl;

    if (!IMG_Init(IMG_INIT_PNG))
        std::cout << "IMG_Init has failed. Error: " << SDL_GetError() << std::endl;

    // Connect to TCP
    auto addr = "127.0.0.1";
    int port = 8080;
    std::cout << "Connecting to: " << addr << ":" << port << std::endl;
    TCPClient tcpClient(io_context);
    tcpClient.start(addr, port, io_context);

    std::thread ioContextThread(runIoContext); // Run io_context.run() in a separate thread

    RenderWindow window("Online Game TCP", 1280, 720);
    SDL_Texture* playerTexture = window.loadTexture("res/imgs/star.png");
    Player player(playerTexture, 5);

    Game game(window, player);
    game.GameLoop(); // main game loop

    tcpClient.stop();

    window.cleanUp();
    SDL_Quit();

    io_context.stop(); // Stop io_context when your program exits

    ioContextThread.join(); // Wait for io_context thread to finish

    return 0;
}
