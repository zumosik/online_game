#pragma once

#include <cstdint>
#include <vector>
#include <cstring>
#include <boost/asio/buffer.hpp>
#include "Types.h"

using bytes = std::vector<std::byte>;

class Buffer {


public:
    explicit Buffer(int maxSize) {
        data = new uint8_t[maxSize];
        size = maxSize;
        index = 0;
    }

    explicit Buffer(const std::vector<char>& vec) : size(vec.size()), index(0) {
        data = new uint8_t[size];
        std::memcpy(data, vec.data(), size);
    }

    ~Buffer() {
        delete[] data;
    }

    uint8_t* GetData();
    size_t GetIndex();

    void Print();
    void AppendByteToFront(uint8_t byte);

    void WriteShort(uint16_t value);
    void WriteChar(uint8_t value);
    void WriteInteger( uint32_t value);
    void WriteDouble(double value);
    void WritePlayer(Player* player);

    uint32_t ReadInteger();
    uint16_t ReadShort();
    uint8_t ReadChar();
    double ReadDouble();
    Player* ReadPlayer();


private:
    uint8_t * data;     // pointer to buffer data
    size_t size;           // size of buffer data (bytes)
    size_t index;          // index of next byte to be read/written
};
