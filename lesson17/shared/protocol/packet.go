package protocol

const (
	HeaderSize  = 6
	MaxBodySize = 4096
)

const (
	PacketLogin  uint16 = 1001
	PacketMove   uint16 = 1002
	PacketAttack uint16 = 1003
)

const (
	PacketMoveNotify uint16 = 2002
)
