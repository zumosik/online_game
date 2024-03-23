//
// Created by a on 23.03.2024.
//

#include "GameObject.h"
#include "TextureManager.h"

GameObject::GameObject(const char *textureSheet, SDL_Renderer *ren) {
    renderer = ren;
    texture = TextureManager::LoadTexture(textureSheet, ren);
}

GameObject::~GameObject() = default;

void GameObject::Update() {
    pos = Vector2f();

    srcRect.h = 32;
    srcRect.w = 32;
    srcRect.x = 0;
    srcRect.y = 0;

    dstRect.h = srcRect.h * 2;
    dstRect.w = srcRect.w * 2;
    dstRect.x = pos.x;
    dstRect.y = pos.y;
}

void GameObject::Render() {
    SDL_RenderCopy(renderer, texture, &srcRect, &dstRect);
}
