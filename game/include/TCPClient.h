#pragma once

#include <string>
#include "Packet.h"
#include "Buffer.h"

#include <boost/asio.hpp>


class TCPClient {
public:
    explicit TCPClient(boost::asio::io_context& io_context, std::string username, uint32_t pin)
            : socket_(io_context), buffer_(1024), username_(username), pin_(pin) {
    }


    void start(const std::string& addr, const int port_num, boost::asio::io_context& io_context);
    void stop();

    void update();

private:
    std::string username_;
    uint32_t pin_;

    boost::asio::ip::tcp::socket socket_;
    std::vector<char> buffer_;
    boost::asio::io_context context;


    void sendBytes(Buffer& buf);
    void sendConnectReq(const ConnectReq& req);
    void handleUpdate();
};