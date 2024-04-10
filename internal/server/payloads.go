package server

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net"
	"online_game/internal/lib/logger/sl"
	"online_game/internal/lib/utils"
	"online_game/internal/models"
	"online_game/internal/packets"
	"strconv"
)

func (s *Server) handleConnectReq(req packets.ConnectReq, conn net.Conn) packets.ConnectResp {
	s.l.Debug("Handling ConnectReq", sl.Attr("username", req.Username), sl.Attr("pin", strconv.Itoa(int(req.Pin))))

	_, exists := s.playerMap[conn]
	if exists { // already connected
		return packets.ConnectResp{OK: false}
	}

	for _, pl := range s.playerMap {
		if pl.Username == req.Username {
			return packets.ConnectResp{OK: false} // someone is already playing
		}
	}

	// Player is not connected:

	var pl models.Player
	pl, playerExists := s.save.Players[req.Username]
	if !playerExists { // create new player
		pl.Pos = models.Vector{X: 1, Y: 1}
		pl.Username = req.Username
		pl.Pin = req.Pin

		// getting unique rnd id
		id := uint16(rand.Intn(65535)) // 65535 - max uint16
		for !s.isUserIDUnique(id) {
			id = uint16(rand.Intn(65535))
		}

		pl.UserID = id

		s.save.Players[pl.Username] = pl
	} else {
		if pl.Pin != req.Pin { // check "password"
			return packets.ConnectResp{OK: false} // pin doesnt match
		}
	}

	for cl := range s.playerMap { // sending to other clients
		s.l.Debug("Sending NewPlayerConnect packet", sl.Attr("player addr", cl.RemoteAddr().String()))
		req := packets.NewPlayerConnect{
			Player: models.PublicPlayer{
				Username: pl.Username,
				UserID:   pl.UserID,
				Pos:      pl.Pos,
			},
		}

		err := s.SendToClient(cl, req, packets.TypeOfPacketNewPlayerConnect)
		if err != nil {
			s.l.Error("Cant send to client about new conn", sl.Err(err), sl.Attr("client who wanted to recieve", cl.RemoteAddr().String()))
		}
	}

	// Token for future requests
	token := utils.GenerateToken()
	pl.Token = token

	s.playerMap[conn] = pl

	s.l.Debug("New player registered", slog.String("username", req.Username), slog.Int("id", int(pl.UserID)))

	// Send to the new player info about all connected players
	var players []models.PublicPlayer
	var i uint8
	for _, player := range s.playerMap {
		if i > s.maxPlayer {
			break
		}

		if player.UserID == pl.UserID {
			continue
		}

		players = append(players, models.PublicPlayer{
			Username: player.Username,
			UserID:   player.UserID,
			Pos:      player.Pos,
		})

		i++
	}

	return packets.ConnectResp{OK: true, AlreadyExists: playerExists,
		Player: models.PublicPlayer{
			Username: pl.Username,
			UserID:   pl.UserID,
			Pos:      pl.Pos,
		},
		Players: players,
		Token:   token,
	}

}

func (s *Server) handlePlayerPosReq(req packets.PlayerPosReq, conn net.Conn) {

	player, exists := s.playerMap[conn]
	if !exists && player.UserID != req.ID {
		return
	}

	if player.Token != req.Token {
		return
	}

	s.l.Debug("Got pos", sl.Attr("x", fmt.Sprintf("%f", req.Vector.X)),
		sl.Attr("y", fmt.Sprintf("%f", req.Vector.Y)))

	player.Pos = req.Vector
	s.playerMap[conn] = player
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
