package main

import (
	"fmt"
	"net"
	"online_game/internal/packets"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var id uint16
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
			return
		}
	}(conn)

	// Set up signal handler to catch SIGINT (Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	go func() {
		<-sigCh
		fmt.Println("\nCtrl+C pressed. Closing connection...")

		req := packets.DisconnectReq{ID: id}
		packet := packets.Packet{
			TypeOfPacket: packets.TypeOfPacketDisconnectReq,
			Payload:      &req,
		}

		b, err := packet.Serialize()
		if err != nil {
			fmt.Println("Error serializing packet:", err)
			return
		}

		sentN, err := conn.Write(b)
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}

		fmt.Println("Sent disconnect")

		fmt.Println("size of sent packet:", sentN)
		if err := conn.Close(); err != nil {
			fmt.Println("Error closing connection:", err)
		}
		os.Exit(0)
	}()

	req := packets.ConnectReq{Username: "user1", Pin: 1488}
	packet := packets.Packet{
		TypeOfPacket: packets.TypeOfPacketConnectReq,
		Payload:      &req,
	}

	b, err := packet.Serialize()
	if err != nil {
		fmt.Println("Error serializing packet:", err)
		return
	}

	sentN, err := conn.Write(b)
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}

	fmt.Println("size of sent packet:", sentN)

	buffer := make([]byte, 1024) // create a buffer to hold incoming data

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	packet, err = packets.Deserialize(buffer[:n])
	if err != nil {
		fmt.Println("Error deserializing packet:", err)
		return
	}

	connResp := packet.Payload.(*packets.ConnectResp)

	if !connResp.OK {
		fmt.Println("!connResp.OK")
		return
	}

	id = connResp.Player.UserID
	vec := connResp.Player.Pos

	for {
		time.Sleep(time.Millisecond * 500)
		start := time.Now()

		vec.X += (1.0 / 3.0)
		vec.Y += (1.0 / 9.0)

		req := packets.PlayerPosReq{ID: id, Vector: vec}
		packet = packets.Packet{
			TypeOfPacket: packets.TypeOfPacketPlayerPosReq,
			Payload:      &req,
		}

		b, err := packet.Serialize()
		if err != nil {
			fmt.Println("Error serializing packet:", err)
			return
		}

		fmt.Println(b)

		_, err = conn.Write(b)
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}

		// n, err := conn.Read(buffer)
		// if err != nil {
		// 	fmt.Println("Error reading from connection:", err)
		// 	return
		// }

		// fmt.Println("size of incoming packet:", n)

		// fmt.Println(buffer[:n])

		// packet, err := server.Deserialize(buffer[:n])
		// if err != nil {
		// 	fmt.Println("Error deserializing packet:", err)
		// 	return
		// }

		end := time.Now().Sub(start)
		fmt.Println(end)
	}
}
