#pragma once

#include <cstdint>
#include <vector>
#include <cstring>
#include <string>
#include <fstream>
#include <iostream>
#include "Types.h"
#include "Buffer.h"


struct ConnectReq {
    uint32_t UsernameLen;
    char Username[20]{};
    uint32_t Pin;

    explicit ConnectReq() {
        std::strncpy(Username, "", sizeof(Username) - 1);
        UsernameLen = static_cast<uint32_t>(std::strlen(Username));
        Pin = 0;
    }

    explicit ConnectReq(const char* username, uint32_t p_pin) {
        std::strncpy(Username, username, sizeof(Username) - 1);
        Username[sizeof(Username) - 1] = '\0'; // Ensure null-terminated
        UsernameLen = static_cast<uint32_t>(std::strlen(Username));
        Pin = p_pin;
    }

    explicit ConnectReq(std::string username, uint32_t p_pin) {
         std::strncpy(Username, username.c_str(), sizeof(Username) - 1);
        Username[sizeof(Username) - 1] = '\0'; // Ensure null-terminated
        UsernameLen = static_cast<uint32_t>(std::strlen(Username));
        Pin = p_pin;
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

struct NewPlayerConnect {
    Player player;

    NewPlayerConnect(): player(){};

    void Write( Buffer & buffer ) const;

    void Read( Buffer & buffer );
};

struct DisconnectReq {
    // TODO
};

struct DisconnectResp {
    // TODO
};


union Payload {
    Empty empty;
    ConnectReq connectReq;
    ConnectResp connectResp;
    NewPlayerConnect newPlayerConnect;
    DisconnectReq  disconnectReq;
    DisconnectResp disconnectResp;
};


enum PacketTypeEnum {EMPTY = 0,
        CONNECT_REQ = 1,
        CONNECT_RESP = 2,
        NEW_PLAYER_CONNECT = 3,
        PACKET_DISCONNECT_REQ = 4,
        PACKET_DISCONNECT_RESP = 5
};


struct Packet
{
    Packet() : packetType(EMPTY), payload{.empty = Empty()} {} // empty packet (for errors os smth like this)
    Packet(PacketTypeEnum type, const ConnectReq& req) : packetType(type), payload{.connectReq = req} {}
    Packet(PacketTypeEnum type, const ConnectResp& resp) : packetType(type), payload{.connectResp = resp} {}
    Packet(PacketTypeEnum type, const NewPlayerConnect& resp) : packetType(type), payload{.newPlayerConnect = resp} {}
    Packet(PacketTypeEnum type, const DisconnectReq& resp) : packetType(type), payload{.disconnectReq = resp} {}
    Packet(PacketTypeEnum type, const DisconnectResp& resp) : packetType(type), payload{.disconnectResp = resp} {}

    PacketTypeEnum packetType;
    Payload payload;


    void Serialize(Buffer & buffer) const; // writes into buffer
    void Deserialize(Buffer & buffer);

};
