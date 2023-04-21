package ethernet

type Ethernet struct {
	ID       string   `json:"id,omitempty"`
	AgentID  string   `json:"agent_id"`
	Name     string   `json:"name"`
	DevName  string   `json:"dev_name"`
	Desc     string   `josn:"desc,omitempty"`
	IPv4     []string `json:"ipv4,omitempty"`
	IPv6     []string `json:"ipv6,omitempty"`
	Hardware string   `json:"hardware"`
}
