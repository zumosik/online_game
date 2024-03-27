#include "TCPClient.hpp"
#include "Packet.hpp"
#include <string>
#include <boost/asio.hpp>
#include <boost/asio/ts/buffer.hpp>
#include <boost/asio/ts/internet.hpp>
#include <thread>
#include "Game.h"

#ifdef _WIN32
#define _WIN32_WINNT 0x0A00
#endif
#define ASIO_STANDALONE

using boost::asio::ip::tcp;


void TCPClient::start(const std::string& addr, const int port_num, boost::asio::io_context& io_context) {
    tcp::resolver resolver(io_context);
    tcp::resolver::results_type endpoints = resolver.resolve(addr, std::to_string(port_num));

    boost::asio::async_connect(socket_, endpoints,
                               [this](const boost::system::error_code& ec, const tcp::endpoint& /*endpoint*/) {
                                   if (!ec) {
                                       std::cout << "Connected to server" << std::endl;

                                       // sending connect msg to server
                                       ConnectReq req("player_cpp");
                                       sendConnectReq(req);



                                       update();
                                   } else {
                                       std::cerr << "Connect error: " << ec.message() << std::endl;
                                       abort();
                                   }
                               });
}

void TCPClient::stop() {
    socket_.close();
}

void TCPClient::update() {
    socket_.async_read_some(boost::asio::buffer(buffer_),
                            [&](std::error_code ec, std::size_t length) {
                                if (!ec) {
                                    Buffer buff(buffer_);
                                    buff.Print();

                                    Packet packet;
                                    packet.Deserialize(buff);
                                    if (packet.packetType == CONNECT_RESP) {
                                        if (packet.payload.connectResp.ok) {
                                            std::cout << "Server registered this connection" << std::endl;

                                            Game::InitializePlayer(&packet.payload.connectResp);
                                        } else {
                                            std::cout << "Server could not register this connection, exiting..." << std::endl;
                                            abort();
                                        }


                                    }



                                    std::this_thread::sleep_for(std::chrono::milliseconds(10)); // 10 ms
                                    update(); // Continue reading
                                } else {
                                    std::cerr << "Read error: " << ec.message() << std::endl;
                                }
                            });
}

void TCPClient::sendBytes(Buffer& buf) {
    async_write(socket_, boost::asio::buffer(buf.GetData(), buf.GetIndex()),
                [](const boost::system::error_code& ec, std::size_t /*bytes_transferred*/) {
                    if (ec) {
                        std::cerr << "Send error: " << ec.message() << std::endl;
                    } else {
                        std::cout << "Sent" << std::endl;
                    }
                });
}

void TCPClient::sendConnectReq(const ConnectReq& req) {
    Packet packet(CONNECT_REQ, req);

    Buffer buf(1024);
    packet.Serialize(buf);

    sendBytes(buf);
}