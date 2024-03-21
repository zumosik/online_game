#include "Entity.hpp"

class Player : public Entity {
public:
    explicit Player(SDL_Texture *pTex, double speed) : Entity(Vector2f(), pTex), velocity(), speed(speed) {}

    Vector2f velocity;
    double speed;
    void Move(); // uses velocity and changes curr pos
};