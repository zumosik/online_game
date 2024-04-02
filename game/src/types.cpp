//
// Created by a on 15.03.2024.
//

#include "Types.hpp"

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


void Player::Write(Buffer &buf, const Player &pl) {
    buf.WriteShort(pl.id);
    auto len = std::strlen(pl.username);
    buf.WriteInteger(static_cast<uint32_t>(len));
    for (int i = 0; i < len; ++i)
        buf.WriteChar( pl.username[i]);

    Vector2f::Write(buf, pl.pos);
}



Player &Player::Read(Buffer &buf) {
    id = buf.ReadShort();
    auto len = buf.ReadInteger();


    for (int i = 0; i < len; ++i)
        username[i] = buf.ReadChar();

    // Null-terminate the username to make it a valid C-string
    username[len] = '\0';
    pos.Read(buf);



    return *this;
}