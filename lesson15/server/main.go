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
		// header (ここではbody size情報)
		header := make([]byte, 4)
		// ReadFull : バッファを全部埋めるまえでRead反復
		_, err := io.ReadFull(conn, header)
		if err != nil {
			return
		}

		// size Check
		size := binary.BigEndian.Uint32(header)

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

		fmt.Println("[Received]: ", string(body))
	}
}
