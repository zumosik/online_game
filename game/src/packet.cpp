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
            // TODO: add more cases idk
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

            break;
        }
        case CONNECT_RESP: {
            ConnectResp v;
            v.Read(buffer);
            payload.connectResp = v;

            break;
        }
        case NEW_PLAYER_CONNECT: {
            NewPlayerConnect v;
            v.Read(buffer);
            payload.newPlayerConnect = v;
            break;
        }
        case PACKET_DISCONNECT_REQ: {
            DisconnectReq v;
            // TODO
            break;
        }
        case PACKET_DISCONNECT_RESP: {
            DisconnectResp v;
            // TODO
            break;

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
    std::cout << "read" << std::endl;
    ok = buffer.ReadChar();
    alreadyExists = buffer.ReadChar();
    player.Read(buffer);


    std::cout << "username: " << player.username << std::endl;
    std::cout << "ok, ae: " << ok << alreadyExists << std::endl;
}

void NewPlayerConnect::Write(Buffer &buffer) const {
    Player::Write(buffer, player);
}

void NewPlayerConnect::Read(Buffer &buffer) {
    player.Read(buffer);
}

