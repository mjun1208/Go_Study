package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
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

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	log.Println("Client Connected:", conn.RemoteAddr())

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
		if size > 4096 {
			return
		}

		// body
		body := make([]byte, int(size))

		_, err = io.ReadFull(conn, body)
		if err != nil {
			return
		}

		switch packetID {
		case 1001:
			fmt.Println("[Login]:", string(body))

		case 1002:
			fmt.Println("[Move]:", string(body))

		case 1003:
			fmt.Println("[Attack]:", string(body))

		default:
			fmt.Println("Unknown Packet ID:", packetID)
		}

		// fmt.Println("[Received]: ", string(body))

		response := append([]byte("[SERVER] "), body...)

		responseHeader := make([]byte, 4)

		binary.BigEndian.PutUint32(
			responseHeader,
			uint32(len(response)),
		)

		packet := append(responseHeader, response...)

		_, err = conn.Write(packet)
		if err != nil {
			log.Println("Write Error:", err)
			return
		}
	}
}
