package agent

type AgentStatusType string

const (
	RegisteredAgentStatus AgentStatusType = "REGISTERED"
	WorkingAgentStatus    AgentStatusType = "WORKING"
	StopedAgentStatus     AgentStatusType = "STOPED"
	FailedAgentStatus     AgentStatusType = "FAILED"
)

type Agent struct {
	ID       string `json:"id,omitempty"`
	Hostname string `json:"hostname"`
	Status   string `json:"status"`
}
