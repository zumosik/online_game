#include <iostream>
#include "Buffer.h"

void Buffer::AppendByteToFront(uint8_t byte) {
    if (index < size) {
        memmove(data + 1, data, index);
        data[0] = byte; // Insert the new byte at the beginning
        ++index;
    }
}

void Buffer::Print() {
    int arraySize = sizeof(data) / sizeof(data[0]);

    std::cout << "Buffer index: " << index << " Buffer size: " << size << std::endl;
    std::cout << "Buffer elements:" << std::endl;
    for (int i = 0; i < arraySize; ++i) {
        std::cout << static_cast<int>(data[i]) << " "; // Cast to int for printing
    }
    std::cout << std::endl;


}

void Buffer::WriteShort(uint16_t value) {
    if (index + 2 <= size) {
        data[index++] = static_cast<uint8_t>((value >> 8) & 0xFF);
        data[index++] = static_cast<uint8_t>(value & 0xFF);
    }
}

void Buffer::WriteChar(uint8_t value) {
    if (index < size) {
        data[index++] = value;
    }
}

void Buffer::WriteInteger(uint32_t value) {
    if (index + 4 <= size) {
        data[index++] = static_cast<uint8_t>((value >> 24) & 0xFF);
        data[index++] = static_cast<uint8_t>((value >> 16) & 0xFF);
        data[index++] = static_cast<uint8_t>((value >> 8) & 0xFF);
        data[index++] = static_cast<uint8_t>(value & 0xFF);
    }
}

void Buffer::WriteDouble(double value) {
    if (index + sizeof(double) <= size) {
        std::memcpy(&data[index], &value, sizeof(double));
        index += sizeof(double);
    }
}

double Buffer::ReadDouble() {
    double value = 0;
    if (index + sizeof(double) <= size) {
        std::memcpy(&value, &data[index], sizeof(double));
        index += sizeof(double);
    }
    return value;
}

uint32_t Buffer::ReadInteger() {
    uint32_t value = 0;
    if (index + 4 <= size) {
        value |= static_cast<uint32_t>(data[index++]) << 24;
        value |= static_cast<uint32_t>(data[index++]) << 16;
        value |= static_cast<uint32_t>(data[index++]) << 8;
        value |= static_cast<uint32_t>(data[index++]);
    }
    return value;
}

uint16_t Buffer::ReadShort() {
    uint16_t value = 0;
    if (index + 2 <= size) {
        value |= static_cast<uint16_t>(data[index++]) << 8;
        value |= static_cast<uint16_t>(data[index++]);
    }
    return value;
}

uint8_t Buffer::ReadChar() {
    uint8_t value = 0;
    if (index < size) {
        value = data[index++];
    }
    return value;
}

uint8_t *Buffer::GetData() {
    return data;
}

size_t Buffer::GetIndex() {
    return index;
}
