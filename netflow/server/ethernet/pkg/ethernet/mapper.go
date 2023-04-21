package ethernet

import "github.com/infEreb/nls/netflow/server/ethernet/src/gen"

func EthernetToProto(e *Ethernet) *gen.Ethernet {
	return &gen.Ethernet{
		Id: e.ID,
		AgentId: e.AgentID,
		Name: e.Name,
		DevName: e.DevName,
		Desc: e.Desc,
		Ipv4: e.IPv4,
		Ipv6: e.IPv6,
		Hardware: e.Hardware,
	}
}

func EthernetFromProto(e *gen.Ethernet) *Ethernet {
	return &Ethernet{
		ID: e.Id,
		AgentID: e.AgentId,
		Name: e.Name,
		DevName: e.DevName,
		Desc: e.Desc,
		IPv4: e.Ipv4,
		IPv6: e.Ipv6,
		Hardware: e.Hardware,
	}
}