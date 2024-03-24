//
// Created by a on 23.03.2024.
//

#ifndef GAME_SDL2_SPRITECOMPONENT_H
#define GAME_SDL2_SPRITECOMPONENT_H

#include "SDL.h"
#include "TransformComponent.h"
#include "TextureManager.h"

class SpriteComponent : public Component {
public:
    SpriteComponent() = default;
    explicit SpriteComponent(const char* path) {
        setTex(path);
    }
    ~SpriteComponent() override {
        SDL_DestroyTexture(tex);
    }

    void setTex(const char* path) {
        tex = TextureManager::LoadTexture(path);
    }

    void init() override {
        transform = &entity->getComponent<TransformComponent>();

        srcRect.x = srcRect.y = 0;
        srcRect.w = transform->w;
        srcRect.h = transform->h;
    }

    void update() override {
        dstRect.x = (int)transform->position.x;
        dstRect.y = (int)transform->position.y;

        dstRect.w = transform->w * transform-> scale;
        dstRect.h = transform->h * transform-> scale;
    }

    void draw() override {
        TextureManager::Draw(tex, srcRect, dstRect);
    }
private:
    TransformComponent *transform;
    SDL_Texture * tex;
    SDL_Rect srcRect, dstRect;
};

#endif //GAME_SDL2_SPRITECOMPONENT_H
