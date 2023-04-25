package packet

type Packet struct {
	ID         string `json:"id,omitempty"`
	AgentID    string `json:"agent_id,omitempty"`
	EthernetID string `json:"ethernet_id,omitempty"`
	Content    string `json:"content,omitempty"`
}