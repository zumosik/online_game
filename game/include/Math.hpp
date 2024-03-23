#pragma once

#include <iostream>

struct Vector2f
{
    Vector2f():x(0.0f), y(0.0f){}
    Vector2f(double p_x,double  p_y): x(p_x), y(p_y) {}

    void print() const {
        std::cout << x << "," << y << std::endl;
    }

    double x,y;
};

struct Vector2int
{
    Vector2int():x(0), y(0){}
    Vector2int(int p_x,int  p_y): x(p_x), y(p_y) {}

    void print() const {
        std::cout << x << "," << y << std::endl;
    }

    void add(Vector2int addVec) {
        x+=addVec.x;
        y+=addVec.y;
    }

    int x,y;
};