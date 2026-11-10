package main

import (
	"errors"
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
			log.Println("Accept Error", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	log.Println("Client Connected:", conn.RemoteAddr())

	// Read
	buf := make([]byte, 1024)

	for {
		// n : 読み込んだbyteの数
		n, err := conn.Read(buf)
		if n > 0 {
			message := buf[:n]
			log.Printf("Received: %q", message)

			response := append([]byte("[SERVER] "), message...)

			// Write
			_, writeErr := conn.Write(response)

			if writeErr != nil {
				log.Println("Write Error:", writeErr)
				return
			}
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Client input closed")
			} else {
				log.Println("Read Error:", err)
			}

			return
		}

	}
}
