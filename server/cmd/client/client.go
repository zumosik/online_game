package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"server/internal/server"
	"syscall"
	"time"
)

func main() {
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
		if err := conn.Close(); err != nil {
			fmt.Println("Error closing connection:", err)
		}
		os.Exit(0)
	}()

	req := server.ConnectReq{Username: "user1"}
	packet := server.Packet{
		TypeOfPacket: server.TypeOfPacketConnectReq,
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

	packet, err = server.Deserialize(buffer[:n])
	if err != nil {
		fmt.Println("Error deserializing packet:", err)
		return
	}

	id := packet.Payload.(*server.ConnectResp).ID

	var vec server.Vector

	vecToAdd := server.Vector{
		X: 1,
		Y: 1,
	}

	for {

		start := time.Now()
		vec.Add(vecToAdd)
		req := server.PlayerPosReq{ID: id, Vector: vec}
		packet = server.Packet{
			TypeOfPacket: server.TypeOfPacketPlayerPosReq,
			Payload:      &req,
		}

		b, err := packet.Serialize()
		if err != nil {
			fmt.Println("Error serializing packet:", err)
			return
		}

		_, err = conn.Write(b)
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		fmt.Println("size of incoming packet:", n)

		fmt.Println(buffer[:n])

		packet, err := server.Deserialize(buffer[:n])
		if err != nil {
			fmt.Println("Error deserializing packet:", err)
			return
		}
		fmt.Println("Received:", *packet.Payload.(*server.PlayerPosResp))

		end := time.Now().Sub(start)
		fmt.Println(end)
	}
}
