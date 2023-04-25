package agent

import "github.com/infEreb/nls/netflow/server/app/agent/src/gen"

func AgentToProto(a *Agent) *gen.Agent {
	return &gen.Agent{
		Id: a.ID,
		Hostname: a.Hostname,
		Status: a.Status,
	}
}

func AgentFromProto(a *gen.Agent) *Agent {
	return &Agent{
		ID: a.Id,
		Hostname: a.Hostname,
		Status: a.Status,
	}
}