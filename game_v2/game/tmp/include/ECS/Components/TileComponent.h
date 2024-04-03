//
// Created by a on 24.03.2024.
//

#ifndef GAME_SDL2_TILECOMPONENT_H
#define GAME_SDL2_TILECOMPONENT_H

#include "ECS/Components.h"

enum Tiles {
    AIR,
    GRASS,
};

class TileComponent : public Component {
public:
    TransformComponent *transform;
    SpriteComponent *sprite;

    SDL_Rect tileRect;
    Tiles tileID;
    char* path;

    TileComponent() = default;

    TileComponent(int x, int y, int w, int h, Tiles id) {
        tileRect.x = x;
        tileRect.y = y;
        tileRect.w = w;
        tileRect.h = h;
        tileID = id;

        switch (id) {
            case GRASS:
                path = "res/imgs/grass.png";
                break;
            default:
                break;
        }
    }

    void init() override {
        entity->addComponent<TransformComponent>(Vector2f(tileRect.x,tileRect.y), 0.0f, tileRect.w, tileRect.h, 1);
        transform = &entity->getComponent<TransformComponent>();

        entity->addComponent<SpriteComponent>(path);
        sprite = &entity->getComponent<SpriteComponent>();
    }
};

#endif //GAME_SDL2_TILECOMPONENT_H
