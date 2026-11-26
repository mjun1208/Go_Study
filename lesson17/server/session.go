package main

import (
	"ServerClient/shared/protocol"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
)

type Session struct {
	ID     int64
	Conn   net.Conn
	sendMu sync.Mutex
}

func (s *Session) Send(packetID uint16, body []byte) error {
	if len(body) > protocol.MaxBodySize {
		return fmt.Errorf("message too large: %d", len(body))
	}

	header := make([]byte, protocol.HeaderSize)

	binary.BigEndian.PutUint32(
		header[0:4],
		uint32(len(body)),
	)

	binary.BigEndian.PutUint16(
		header[4:6],
		packetID,
	)

	packet := append(header, body...)

	s.sendMu.Lock()
	defer s.sendMu.Unlock()

	_, err := s.Conn.Write(packet)
	return err
}
