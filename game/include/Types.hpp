#pragma once

#include <iostream>
#include "Buffer.hpp"

struct Vector2f
{
    Vector2f();
    Vector2f(double p_x,double  p_y);
    Vector2f(int p_x,int  p_y);


    Vector2f& Add(const Vector2f& vec);
    Vector2f& Subtract(const Vector2f& vec);
    Vector2f& Multiply(const Vector2f& vec);
    Vector2f& Divide(const Vector2f& vec);

    friend Vector2f&   operator+(Vector2f& v1, const Vector2f& v2);
    friend Vector2f&   operator-(Vector2f& v1, const Vector2f& v2);
    friend Vector2f&   operator*(Vector2f& v1, const Vector2f& v2);
    friend Vector2f&   operator/(Vector2f& v1, const Vector2f& v2);

    Vector2f& operator += (const Vector2f& vec);
    Vector2f& operator -= (const Vector2f& vec);
    Vector2f& operator *= (const Vector2f& vec);
    Vector2f& operator /= (const Vector2f& vec);

    Vector2f& operator*(const float & i);
    Vector2f& Zero();

    static void Write( Buffer & buf, const Vector2f& vec);
    Vector2f& Read( Buffer & buf);

    friend std::ostream& operator << (std::ostream& ostream, const Vector2f& vec);

    double x,y;
};

struct Player {
    char username[20]{};
    uint16_t id;
    Vector2f pos;

    Player() :username(""), id(0), pos(Vector2f())  {};

    static void Write( Buffer & buf, const Player& pl);
    Player& Read( Buffer & buf);

//    uint32_t pin;
};