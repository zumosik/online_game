package server

import (
	"log/slog"
	"math/rand"
	"net"
	"server/internal/models"
)

func (s *Server) handleConnectReq(req ConnectReq, conn net.Conn) ConnectResp {
	_, exists := s.playerMap[conn]
	if exists { // already connected
		return ConnectResp{OK: false}
	}

	for _, pl := range s.playerMap {
		if pl.Username == req.Username {
			return ConnectResp{OK: false} // someone is already playing
		}
	}

	// Player is not connected:

	var pl models.Player
	pl, playerExists := s.save.Players[req.Username]
	if !playerExists {
		pl.Pos = models.Vector{X: 1, Y: 1}
		pl.Username = req.Username

		// getting unique rnd id
		id := uint16(rand.Intn(65535))
		for !s.isUserIDUnique(id) {
			id = uint16(rand.Intn(65535))
		}

		pl.UserID = id

		// we don't need to save player here because it will be saves on shutdown
	}

	s.playerMap[conn] = pl

	s.l.Debug("New player registered", slog.String("username", req.Username), slog.Int("id", int(pl.UserID)))

	return ConnectResp{OK: true, AlreadyExists: playerExists, Player: pl}

}

func (s *Server) handlePlayerPosReq(req PlayerPosReq, conn net.Conn) {

	player, exists := s.playerMap[conn]
	if !exists && player.UserID != req.ID {
		return
	}

	player.Pos = req.Vector
	s.playerMap[conn] = player
}

func (s *Server) handleDisconnect(req DisconnectReq, conn net.Conn) DisconnectResp {
	s.connClose(conn)

	var resp DisconnectResp
	resp.OK = true

	return resp
}

func (s *Server) isUserIDUnique(userID uint16) bool {
	// Iterate over the playerMap and check if the UserID already exists
	for _, player := range s.playerMap {
		if player.UserID == userID {
			return false // UserID is not unique
		}
	}
	return true // UserID is unique
}
