//
// Created by a on 15.03.2024.
//

#include "Math.hpp"

Vector2f::Vector2f() {x = 0.0f; y = 0.0f;}

Vector2f::Vector2f(int p_x, int p_y) {
    this->x = p_x;
    this->y = p_x;
}



Vector2f::Vector2f(double p_x, double p_y) {
    this->x = p_x;
    this->y = p_y;
}

Vector2f &Vector2f::Add(const Vector2f &vec) {
    this->x += vec.x;
    this->y += vec.y;

    return *this;
}

Vector2f &Vector2f::Subtract(const Vector2f &vec) {
    this->x -= vec.x;
    this->y -= vec.y;
    return *this;


}

Vector2f &Vector2f::Multiply(const Vector2f &vec) {
    this->x *= vec.x;
    this->y *= vec.y;
    return *this;

}

Vector2f &Vector2f::Divide(const Vector2f &vec) {
    this->x /= vec.x;
    this->y /= vec.y;
    return *this;

}

Vector2f &operator+(Vector2f &v1, const Vector2f &v2) {
    return v1.Add(v2);
}

Vector2f &operator-(Vector2f &v1, const Vector2f &v2) {
    return v1.Subtract(v2);

}

Vector2f &operator*(Vector2f &v1, const Vector2f &v2) {
    return v1.Multiply(v2);

}

Vector2f &operator/(Vector2f &v1, const Vector2f &v2) {
    return v1.Divide(v2);

}

Vector2f &Vector2f::operator+=(const Vector2f &vec) {
    this->Add(vec);
}

Vector2f &Vector2f::operator-=(const Vector2f &vec) {
    this->Subtract(vec);

}

Vector2f &Vector2f::operator*=(const Vector2f &vec) {
    this->Multiply(vec);

}

Vector2f &Vector2f::operator/=(const Vector2f &vec) {
    this->Divide(vec);

}

std::ostream &operator<<(std::ostream &ostream, const Vector2f &vec) {
    ostream << "Vec2f(" << vec.x << ", " << vec.y << ")";
    return ostream;
}


Vector2f &Vector2f::operator*(const float &i) {
    this->x *= i;
    this->y *= i;

    return *this;
}

Vector2f &Vector2f::Zero() {
    this->x = 0;
    this->y = 0;

    return *this;
}

void Vector2f::Write(Buffer &buf, const Vector2f &vec) {
    buf.WriteDouble(vec.x);
    buf.WriteDouble(vec.y);
}

Vector2f &Vector2f::Read(Buffer &buf) {
    this->x = buf.ReadDouble();
    this->y = buf.ReadDouble();

    return *this;
}


