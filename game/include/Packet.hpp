#pragma once

#include <cstdint>
#include <vector>
#include <cstring>
#include <fstream>
#include <iostream>
#include "Math.hpp"
#include "Buffer.hpp"
#include "Types.h"


struct ConnectReq {
    uint32_t UsernameLen;
    char Username[20]{};

    explicit ConnectReq() {
        std::strncpy(Username, "", sizeof(Username) - 1);
        UsernameLen = static_cast<uint32_t>(std::strlen(Username));
    }

    explicit ConnectReq(const char* username) {
        std::strncpy(Username, username, sizeof(Username) - 1);
        Username[sizeof(Username) - 1] = '\0'; // Ensure null-terminated
        UsernameLen = static_cast<uint32_t>(std::strlen(Username));
    }

    void Write( Buffer & buffer ) const;

    void Read( Buffer & buffer );
};

struct ConnectResp {
    bool ok;
    bool alreadyExists;
    Player player;

    ConnectResp(): ok(false), alreadyExists(false), player(Player()) {};

    void Write( Buffer & buffer ) const;

    void Read( Buffer & buffer );
};

struct Empty {

};


union Payload {
    Empty empty;
    ConnectReq connectReq;
    ConnectResp connectResp;
};

enum PacketTypeEnum {EMPTY = 0,  CONNECT_REQ = 1, CONNECT_RESP = 2 };

struct Packet
{
    Packet() : packetType(EMPTY), payload{.empty = Empty()} {} // empty packet (for errors os smth like this)
    Packet(PacketTypeEnum type, const ConnectReq& req) : packetType(type), payload{.connectReq = req} {}
    Packet(PacketTypeEnum type, const ConnectResp& resp) : packetType(type), payload{.connectResp = resp} {}

    PacketTypeEnum packetType;
    Payload payload;


    void Serialize(Buffer & buffer) const; // writes into buffer
    void Deserialize(Buffer & buffer);

};
