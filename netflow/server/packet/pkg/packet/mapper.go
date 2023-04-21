package packet

import (
	"encoding/base64"

	"github.com/infEreb/nls/netflow/server/packet/src/gen"
)

func PacketToProto(p *Packet) *gen.Packet {
	s, _ := base64.RawStdEncoding.DecodeString(p.Content)
	return &gen.Packet{
		Id: p.ID,
		AgentId: p.AgentID,
		EthernetId: p.EthernetID,
		Content: s,
	}
}

func PacketFromProto(p *gen.Packet) *Packet {
	s := base64.RawStdEncoding.EncodeToString(p.Content)
	return &Packet{
		ID: p.Id,
		AgentID: p.AgentId,
		EthernetID: p.EthernetId,
		Content: s,
	}
}