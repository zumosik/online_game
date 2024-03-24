cmake_minimum_required(VERSION 3.27)
project(game_sdl2)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_MODULE_PATH ${CMAKE_SOURCE_DIR}/cmake_modules)

set(SDL2_PATH "C:/SDL2")
set(SDL2_IMAGE_PATH "C:/SDL2_image")

find_package(Boost 1.84.0 REQUIRED)
find_package(SDL2 REQUIRED)

include(cmake_modules/FindSDL2_image.cmake)
include_directories(${SDL2_IMAGE_INCLUDE_DIRS})
include_directories(${SDL2_INCLUDE_DIR})
include_directories(${Boost_INCLUDE_DIRS})
include_directories(include)


add_executable(game_sdl2
        src/main.cpp
        src/math.cpp
        src/packet.cpp
        src/tcpClient.cpp
        src/buffer.cpp
        src/Game.cpp
        src/textureManager.cpp
        src/Map.cpp
        src/Collision.cpp
)
target_link_libraries(${PROJECT_NAME} ${SDL2_LIBRARY} ${SDL2_IMAGE_LIBRARIES} ${Boost_LIBRARIES} ws2_32)
