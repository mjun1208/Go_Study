package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Server struct {
	Connections atomic.Int64
}

func (server *Server) Connect() {
	server.Connections.Add(1)
}

func (server *Server) Disconnect() {
	server.Connections.Add(-1)
}

func (server *Server) GetConnectionCount() int64 {
	return server.Connections.Load()
}

func main() {
	server := &Server{}
	server.Connections.Store(0)

	var wg sync.WaitGroup
	wg.Add(130)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			server.Connect()
		}()
	}

	for i := 0; i < 30; i++ {
		go func() {
			defer wg.Done()

			server.Disconnect()
		}()
	}

	wg.Wait()

	fmt.Println("Connections : ", server.GetConnectionCount())
}
