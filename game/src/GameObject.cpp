//
// Created by a on 23.03.2024.
//

#include "GameObject.h"
#include "TextureManager.h"
#include "Game.h"

GameObject::GameObject(const char *textureSheet, Vector2int position) {
    texture = TextureManager::LoadTexture(textureSheet);

    pos = position;
}

GameObject::~GameObject() = default;

void GameObject::Update() {
    pos.add(Vector2int(1,1));

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
    SDL_RenderCopy(Game::renderer, texture, &srcRect, &dstRect);
}
