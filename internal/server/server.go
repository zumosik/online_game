package server

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"online_game/internal/lib/logger/sl"
	"online_game/internal/models"
	"online_game/internal/packets"
	"online_game/internal/saver"
	"os"
	"os/signal"
	"syscall"
)

var (
	ErrInvalidType = errors.New("invalid type of packet")
)

type Message struct {
	from   net.Conn
	packet packets.Packet
}

type Config struct {
	Addr        string
	Logger      *slog.Logger
	MaxReadSize uint32

	PathToSave string
}

type Server struct {
	listenAddr string
	ln         net.Listener
	l          *slog.Logger

	maxReadSize uint32
	maxPlayer   uint8

	quitCh chan struct{}
	msgCh  chan Message

	save     *saver.Save
	savePath string

	playerMap map[net.Conn]models.Player
}

func New(cfg *Config) *Server {
	f, err := os.Open(cfg.PathToSave)
	if err != nil {
		panic(err)
	}

	//var save *saver.Save - doesnt work
	save := &saver.Save{}

	cfg.Logger.Info("Loading all data from file", sl.Attr("path", cfg.PathToSave))
	err = save.ReadFromFile(f)
	if err != nil {
		panic(err)
	}

	save.Print(cfg.Logger)

	return &Server{
		listenAddr: cfg.Addr,
		l:          cfg.Logger,

		maxReadSize: cfg.MaxReadSize,
		maxPlayer:   10,

		quitCh: make(chan struct{}),
		msgCh:  make(chan Message, 10),

		// saves
		save:     save,
		savePath: cfg.PathToSave,

		playerMap: make(map[net.Conn]models.Player),
	}
}

func (s *Server) MustStart() {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		s.l.Error("cant listen", sl.Err(err))
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			s.l.Error("cant close listener", sl.Err(err))
		}
	}(ln)
	s.ln = ln

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalCh
		s.l.Info("Received termination signal. Shutting down server...")
		s.shutdown()
	}()

	s.l.Info("Starting server!")

	go s.acceptLoop()
	go s.msgLoop()

	<-s.quitCh

	close(s.msgCh)

	err = s.SaveToSaver()
	if err != nil {
		s.l.Error("cant save", sl.Err(err))
	}
}

func (s *Server) acceptLoop() {
	for {
		select {
		case <-s.quitCh:
			return // Exit the accept loop if shutdown signal received
		default:
			conn, err := s.ln.Accept()
			if err != nil {
				s.l.Error("cant accept new conn", sl.Err(err))
				continue
			}

			s.l.Info("new connection", sl.Attr("addr", conn.RemoteAddr().String()))

			go s.readLoop(conn)
		}
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer s.connClose(conn)
	buf := make([]byte, s.maxReadSize)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			if errors.Is(err, io.EOF) { // closed conn
				s.l.Debug("conn closed", sl.Attr("addr", conn.RemoteAddr().String()))
				return
			}
			s.l.Error("cant read", sl.Err(err))
		}

		// check if something was read
		if n == 0 {
			continue
		} else {
			s.l.Debug("read", sl.Attr("bytes", fmt.Sprint(n)))
		}

		s.l.Debug("received data", sl.Attr("data", fmt.Sprintf("%v", buf[:n])))

		packet, err := packets.DeserializePacket(buf)
		if err != nil {
			s.l.Error("cant deserialize packet", sl.Err(err))
			continue
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
		case packets.TypeOfPacketConnectReq:
			req := msg.packet.Payload.(packets.ConnectReq)
			resp := s.handleConnectReq(req, msg.from)
			err := s.SendToClient(msg.from, resp, packets.TypeOfPacketConnectResp)
			if err != nil {
				s.l.Error("cant send to client", sl.Err(err))
				continue
			}
		case packets.TypeOfPacketPlayerPosReq:
			req := msg.packet.Payload.(packets.PlayerPosReq)
			s.handlePlayerPosReq(req, msg.from)
		default:
			// idk
		}
	}
}

func (s *Server) SendToClient(conn net.Conn, payload interface{}, typeOfPacket byte) error {
	packet := packets.Packet{
		TypeOfPacket: typeOfPacket,
		Payload:      payload,
	}

	data, err := packets.SerializePacket(packet)
	if err != nil {
		return err
	}

	_, err = conn.Write(data)

	s.l.Debug("sent data", sl.Attr("data", fmt.Sprintf("%v", data)))

	return err
}

func (s *Server) connClose(conn net.Conn) {
	pl, exists := s.playerMap[conn]
	if !exists {
		return
	}

	s.save.Players[pl.Username] = pl
	s.l.Debug("Saved player", sl.Attr("username", pl.Username), sl.Attr("id", fmt.Sprint(pl.UserID)))
	delete(s.playerMap, conn) // deleting player from map

	err := conn.Close()
	if err != nil {
		s.l.Error("cant close conn", sl.Err(err))
	}

}

func (s *Server) SaveToSaver() error {
	s.l.Info("Saving all data to file", sl.Attr("path", s.savePath))
	f, err := os.Create(s.savePath)
	if err != nil {
		return err
	}

	for c := range s.playerMap {
		s.connClose(c)
	}

	err = s.save.WriteToFile(f)
	return err
}

func (s *Server) shutdown() {
	close(s.quitCh)
}
