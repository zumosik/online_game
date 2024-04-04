#pragma once

#include <string>
#include "Packet.h"
#include "Buffer.h"

#include <boost/asio.hpp>


class TCPClient {
public:
    explicit TCPClient(boost::asio::io_context& io_context)
            : socket_(io_context), buffer_(1024){
    }


    void start(const std::string& addr, const int port_num, boost::asio::io_context& io_context);
    void stop();

    void update();

private:
    boost::asio::ip::tcp::socket socket_;
    std::vector<char> buffer_;
    boost::asio::io_context context;
    std::chrono::steady_clock::time_point start_time;


    void sendBytes(Buffer& buf);
    void sendConnectReq(const ConnectReq& req);
    void handleUpdate();
};