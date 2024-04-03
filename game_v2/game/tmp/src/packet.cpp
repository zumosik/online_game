#include "Packet.hpp"
#include "Buffer.hpp"



void Packet::Serialize(Buffer &buffer) const {
    switch (packetType) {
        case CONNECT_REQ:
            payload.connectReq.Write(buffer);
            break;
        case CONNECT_RESP:
            payload.connectResp.Write(buffer);
            break;
        default:
            break;
    }

    buffer.AppendByteToFront(static_cast<uint8_t>(packetType));
}

void Packet::Deserialize(Buffer &buffer) {
//    packetType = buffer.ReadChar();
    char t = buffer.ReadChar();
    packetType = static_cast<PacketTypeEnum>(t);


    switch (packetType) {
        case CONNECT_REQ: {
            ConnectReq v;
            v.Read(buffer);
            payload.connectReq = v;
        }
        case CONNECT_RESP: {
            ConnectResp v;
            v.Read(buffer);
            payload.connectResp = v;
        }
        default:
            break;
    }
}

void ConnectReq::Write(Buffer &buffer) const {
    buffer.WriteInteger( UsernameLen);

    for (int i = 0; i < UsernameLen; ++i)
        buffer.WriteChar( Username[i]);

}
void ConnectReq::Read(Buffer &buffer) {
    UsernameLen = buffer.ReadInteger();
    for (int i = 0; i < UsernameLen; ++i)
        Username[i] = buffer.ReadChar();
}

void ConnectResp::Write(Buffer &buffer) const {
    buffer.WriteChar( ok);
    buffer.WriteChar( alreadyExists);
    Player::Write(buffer, player);
}
void ConnectResp::Read(Buffer &buffer) {
    ok = buffer.ReadChar();
    alreadyExists = buffer.ReadChar();
    player.Read(buffer);
}