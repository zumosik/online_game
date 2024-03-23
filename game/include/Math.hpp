#pragma once

#include <iostream>

struct Vector2f
{
    Vector2f();
    Vector2f(double p_x,double  p_y);


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

    friend std::ostream& operator << (std::ostream& ostream, const Vector2f& vec);

    double x,y;
};
