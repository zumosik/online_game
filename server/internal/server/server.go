package server

import (
	"errors"
	"io"
	"log/slog"
	"net"
	"server/utils"
)

var (
	ErrInvalidType = errors.New("invalid type of packet")
)

type Message struct {
	from   net.Conn
	packet Packet
}

type Config struct {
	Addr        string
	Logger      *slog.Logger
	MaxReadSize uint32
}

type Server struct {
	listenAddr string
	ln         net.Listener
	l          *slog.Logger

	maxReadSize uint32
	quitCh      chan struct{}
	msgCh       chan Message

	playerMap map[net.Conn]Player
}

func New(cfg *Config) *Server {
	return &Server{
		listenAddr: cfg.Addr,
		l:          cfg.Logger,

		maxReadSize: cfg.MaxReadSize,
		quitCh:      make(chan struct{}),
		msgCh:       make(chan Message, 10),

		playerMap: make(map[net.Conn]Player),
	}
}

func (s *Server) MustStart() {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		s.l.Error("cant listen", utils.WrapErr(err))
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			s.l.Error("cant close listener", utils.WrapErr(err))
		}
	}(ln)
	s.ln = ln

	s.l.Info("Starting server!")

	go s.acceptLoop()
	go s.msgLoop()

	<-s.quitCh

	close(s.msgCh)

}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			s.l.Error("cant accept new conn", utils.WrapErr(err))
			continue
		}

		s.l.Info("new connection", utils.Wrap("addr", conn.RemoteAddr().String()))

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer s.connClose(conn)
	buf := make([]byte, s.maxReadSize)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) { // closing conn
				s.l.Debug("conn closed", utils.Wrap("addr", conn.RemoteAddr().String()))
				return
			}
			s.l.Error("cant read", utils.WrapErr(err))
			return // usually when client conn is closed

		}

		_ = n

		packet, err := Deserialize(buf[:n])
		if err != nil {
			s.l.Error("cant deserialize packet", utils.WrapErr(err))
		}

		s.msgCh <- Message{
			from:   conn,
			packet: packet,
		}
	}
}

func (s *Server) msgLoop() {
	for msg := range s.msgCh {
		if msg.packet.Payload == nil {
			s.l.Debug("payload is nil, continuing with next package")
			continue
		}
		switch msg.packet.TypeOfPacket {
		case TypeOfPacketConnectReq:
			req := msg.packet.Payload.(*ConnectReq)
			resp := s.handleConnectReq(*req, msg.from)
			err := s.SendToClient(msg.from, &resp)
			if err != nil {
				s.l.Error("cant send to client", utils.WrapErr(err))
				continue
			}
		case TypeOfPacketPlayerPosReq:
			req := msg.packet.Payload.(*PlayerPosReq)
			resp := s.handlePlayerPosReq(*req, msg.from)
			err := s.SendToClient(msg.from, &resp)
			if err != nil {
				s.l.Error("cant send to client", utils.WrapErr(err))
				continue
			}
		}
	}
}

func (s *Server) SendToClient(conn net.Conn, payload Payload) error {
	var typeOfPacket uint8

	switch payload.(type) {
	case *ConnectResp:
		typeOfPacket = TypeOfPacketConnectResp
	case *PlayerPosResp:
		typeOfPacket = TypeOfPacketPlayerPosResp
	default:
		return ErrInvalidType
	}

	packet := Packet{
		TypeOfPacket: typeOfPacket,
		Payload:      payload,
	}

	b, err := packet.Serialize()
	if err != nil {
		return err
	}

	_, err = conn.Write(b)
	return err
}

func (s *Server) connClose(conn net.Conn) {
	delete(s.playerMap, conn) // deleting player from map

	err := conn.Close()
	if err != nil {
		s.l.Error("cant close conn", utils.WrapErr(err))
	}
}
