//
// Created by a on 23.03.2024.
//

#ifndef GAME_SDL2_SPRITETEXTCOMPONENT_H
#define GAME_SDL2_SPRITETEXTCOMPONENT_H

#include "SDL.h"
#include "TransformComponent.h"
#include "TextureManager.h"

class SpriteTextComponent : public Component {
public:
    SpriteTextComponent() = default;
    SpriteTextComponent(const char* path_to_ttf, int size, SDL_Color color, const char* text) {
        setTex(path_to_ttf, size, color, text);
    }
    ~SpriteTextComponent() override {
        SDL_DestroyTexture(tex);
    }

    void setTex(const char* path_to_ttf, int size, SDL_Color color, const char* text) {
        tex = TextureManager::LoadTTFTexture(path_to_ttf, size, color, text, srcRect.w, srcRect.h);
    }

    void init() override {
        transform = &entity->getComponent<TransformComponent>();

        srcRect.x = srcRect.y = 0;
    }

    void update() override {
        dstRect.x = static_cast<int>(transform->position.x);
        dstRect.y = static_cast<int>(transform->position.y);

        dstRect.w = srcRect.w * transform-> scale;
        dstRect.h = srcRect.h * transform-> scale;
    }

    void draw() override {
        TextureManager::Draw(tex, srcRect, dstRect);
    }
private:
    TransformComponent *transform;
    SDL_Texture * tex;
    SDL_Rect srcRect, dstRect;
};

#endif //GAME_SDL2_SPRITETEXTCOMPONENT_H
