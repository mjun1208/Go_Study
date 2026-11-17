package main

import (
	"ServerClient/shared/protocol"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	log.Println("Connected to Server!")

	err = sendPacket(conn, protocol.PacketLogin, "Hello")
	if err != nil {
		log.Println("Error:", err)
		return
	}

	err = sendPacket(conn, protocol.PacketLogin, "World")
	if err != nil {
		log.Println("Error:", err)
		return
	}

	err = sendPacket(conn, protocol.PacketMove, "Move")
	if err != nil {
		log.Println("Error:", err)
		return
	}

	err = sendPacket(conn, protocol.PacketAttack, "Attack")
	if err != nil {
		log.Println("Error:", err)
		return
	}
}

func sendPacket(conn net.Conn, packetID uint16, message string) error {
	body := []byte(message)

	header := make([]byte, protocol.HeaderSize)
	binary.BigEndian.PutUint32(header[0:4], uint32(len(body)))
	binary.BigEndian.PutUint16(header[4:6], packetID)

	packet := append(header, body...)

	_, err := conn.Write(packet)
	if err != nil {
		// %w : エラーを新しいエラーに包む
		return fmt.Errorf("Write Error: %w", err)
	}

	err = receivePacket(conn)

	return err
}

func receivePacket(conn net.Conn) error {
	responseHeader := make([]byte, protocol.HeaderSize)

	_, err := io.ReadFull(conn, responseHeader)
	if err != nil {
		log.Println("Header Error:", err)
		return err
	}

	size := binary.BigEndian.Uint32(responseHeader)

	if size > protocol.MaxBodySize {
		return fmt.Errorf("Response Too Large: %d", size)
	}

	responseBody := make([]byte, int(size))
	_, err = io.ReadFull(conn, responseBody)
	if err != nil {
		log.Println("Body Error:", err)
		return err
	}

	log.Println("[Received]:", string(responseBody))

	return err
}
