package ethernet

type Ethernet struct {
	ID       string   `json:"id,omitempty"`
	AgentID  string   `json:"agent_id,omitempty"`
	Name     string   `json:"name,omitempty"`
	DevName  string   `json:"dev_name,omitempty"`
	Desc     string   `json:"desc,omitempty"`
	IPv4     []string `json:"ipv4,omitempty"`
	IPv6     []string `json:"ipv6,omitempty"`
	Hardware string   `json:"hardware,omitempty"`
}
