package main

import (
	"log"
	"net"
	"sync"
)

type Room struct {
	mu       sync.Mutex
	sessions map[int64]*Session
}

func NewRoom() *Room {
	return &Room{
		sessions: make(map[int64]*Session),
	}
}

func (r *Room) Join(session *Session) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.sessions[session.ID] = session
}

func (r *Room) JoinWithConn(id int64, conn net.Conn) *Session {
	session := &Session{
		ID:   id,
		Conn: conn,
	}

	r.Join(session)

	return session
}

func (r *Room) Leave(id int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.sessions, id)
}

func (r *Room) Count() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	return len(r.sessions)
}

func (r *Room) BroadcastExcept(
	senderID int64,
	packetID uint16,
	body []byte,
) {
	r.mu.Lock()

	recipients := make([]*Session, 0, len(r.sessions))

	for id, session := range r.sessions {
		if id != senderID {
			recipients = append(recipients, session)
		}
	}

	r.mu.Unlock()

	for _, session := range recipients {
		err := session.Send(packetID, body)
		if err != nil {
			log.Printf("Send to session %d failed: %v", session.ID, err)
		}

		log.Printf(
			"[Broadcast SENT] Session=%d, PacketID=%d",
			session.ID, packetID,
		)
	}
}
