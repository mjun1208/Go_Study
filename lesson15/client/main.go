package main

import (
	"encoding/binary"
	"fmt"
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

	err = sendPacket(conn, "Hello")
	if err != nil {
		log.Println("Write Error:", err)
		return
	}

	err = sendPacket(conn, "World")
	if err != nil {
		log.Println("Write Error:", err)
		return
	}
}

func sendPacket(conn net.Conn, message string) error {
	body := []byte(message)

	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(len(body)))

	packet := append(header, body...)

	_, err := conn.Write(packet)
	fmt.Println("[Send]: ", string(body))
	return err
}
