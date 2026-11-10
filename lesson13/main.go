package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Listen Error", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server Started!")

	for {
		// listener.Accept : 新しい連結が来るまで待つ
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept Error", err)
			continue
		}

		// goroutineじゃないと待つ可能性があり
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client Connected:", conn.RemoteAddr())

	time.Sleep(3 * time.Second)

	err := conn.Close()

	if err != nil {
		fmt.Println("Close Error:", err)
		return
	}

	fmt.Println("Client Disconnected:", conn.RemoteAddr())
}
