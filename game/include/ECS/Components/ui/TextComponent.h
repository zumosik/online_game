//
// Created by a on 28.03.2024.
//

#ifndef GAME_SDL2_TEXTCOMPONENT_H
#define GAME_SDL2_TEXTCOMPONENT_H

#include "ECS/Components.h"
#include "SDL_ttf.h"

class TextComponent : public Component {
public:

    TextComponent() = default;

    TextComponent(const char *file_path, int size, SDL_Color color)   {
        std::cout << "eu000" << std::endl;
        TTF_Font  * font = TTF_OpenFont(file_path, size);
        SDL_Surface * textSurf = TTF_RenderText_Solid(font, "hello world", color);
        textTexture = SDL_CreateTextureFromSurface(Game::renderer, textSurf);
        SDL_FreeSurface(textSurf);
        TTF_CloseFont(font);
        std::cout << "eu001" << std::endl;

    }

    ~TextComponent() {
        SDL_DestroyTexture(textTexture);
    }

    void init() override {
        std::cout << "eu001____" << std::endl;

        if (!entity->hasComponent<TransformComponent>()) {
            entity->addComponent<TransformComponent>();
        }

        transform = &entity->getComponent<TransformComponent>();

        srcRect.x = srcRect.y = 0;
        srcRect.w = transform->w;
        srcRect.h = transform->h;

        std::cout << "eu001____1" << std::endl;

    }

    void update() override {
        std::cout << "eu000_update_0" << std::endl;

        dstRect.x = static_cast<int>(transform->position.x);
        dstRect.y = static_cast<int>(transform->position.y);

        dstRect.w = transform->w * transform-> scale;
        dstRect.h = transform->h * transform-> scale;
        std::cout << "eu000_update_1" << std::endl;

    }

    void draw() override {
        std::cout << "eu000_draw_0" << std::endl;

        TextureManager::Draw(textTexture, srcRect, dstRect);

        std::cout << "eu000_draw_1" << std::endl;

    }
private:
    SDL_Texture * textTexture;
    TransformComponent * transform;
    SDL_Rect srcRect, dstRect;
};

#endif //GAME_SDL2_TEXTCOMPONENT_H
