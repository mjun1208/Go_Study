package main

import (
	"ServerClient/shared/protocol"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync/atomic"
)

type PacketHandler func([]byte) []byte

var handlers = map[uint16]PacketHandler{
	protocol.PacketLogin:  handleLogin,
	protocol.PacketMove:   handleMove,
	protocol.PacketAttack: handleAttack,
}

func main() {
	var nextSessionID atomic.Int64

	room := NewRoom()

	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer listener.Close()

	log.Println("Server Started!")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Accept Error : ", err)
			continue
		}

		sessionID := nextSessionID.Add(1)

		go handleClient(conn, room, sessionID)
	}
}

func handleClient(conn net.Conn, room *Room, sessionID int64) {
	session := room.JoinWithConn(sessionID, conn)

	defer func() {
		room.Leave(sessionID)
		conn.Close()
		log.Printf("Player%d DisConnected", sessionID)
	}()

	log.Printf("Player%d Connected (Room Count: %d)", sessionID, room.Count())

	for {
		// header
		header := make([]byte, 6)

		_, err := io.ReadFull(conn, header)
		if err != nil {
			return
		}

		// size Check
		size := binary.BigEndian.Uint32(header[0:4])
		packetID := binary.BigEndian.Uint16(header[4:6])

		// 検証
		if size > protocol.MaxBodySize {
			return
		}

		// body
		body := make([]byte, int(size))

		_, err = io.ReadFull(conn, body)
		if err != nil {
			return
		}

		response := dispatchPacket(packetID, body)

		err = session.Send(packetID, response)
		if err != nil {
			log.Println("Response Error:", err)
			return
		}

		if packetID == protocol.PacketMove {
			message := []byte(fmt.Sprintf("Session%d moved: %s", sessionID, body))

			room.BroadcastExcept(sessionID, protocol.PacketMoveNotify, message)
		}
	}
}

func dispatchPacket(packetID uint16, body []byte) []byte {
	handler, ok := handlers[packetID]

	if !ok {
		return []byte("Unknown Packet")
	}

	return handler(body)
}

func handleLogin(body []byte) []byte {
	log.Println("[Login]:", string(body))
	return []byte("Login Success")
}

func handleMove(body []byte) []byte {
	log.Println("[Move]:", string(body))
	return []byte("Move Success")
}

func handleAttack(body []byte) []byte {
	log.Println("[Attack]:", string(body))
	return []byte("Attack Success")
}
