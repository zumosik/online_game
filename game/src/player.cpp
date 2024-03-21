#include "Player.hpp"

void Player::Move() {
    Vector2f position = getPos();

    position.x += velocity.x * speed;
    position.y += velocity.y * speed;

    setPos(position);
}