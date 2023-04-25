package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/infEreb/nls/netflow/server/app/agent/pkg/agent"
	"github.com/infEreb/nls/netflow/server/pkg/discovery"
	"github.com/infEreb/nls/netflow/server/pkg/netflow"
	"github.com/infEreb/nls/netflow/server/pkg/serror"
)

type ServiceNameGateway string

const agentServiceName ServiceNameGateway = "agent"

type AgentGatewayHTTP struct {
	registry discovery.Registry
}

func NewAgentGatewayHTTP(registry discovery.Registry) *AgentGatewayHTTP {
	return &AgentGatewayHTTP{
		registry: registry,
	}
}

func (g *AgentGatewayHTTP) GetAgent(ctx context.Context, id string) (*agent.Agent, error) {
	addrs, err := g.registry.ServiceAddresses(ctx, string(agentServiceName))
	if err != nil {
		return nil, err
	}

	url := "http://" + addrs[rand.Intn(len(addrs))] + netflow.ROOT + agent.ROOT + "/" + id
	log.Println("Calling agent service. Req: GET " + url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		var err string
		if err := json.NewDecoder(resp.Body).Decode(&err); err != nil {
			log.Println(err)
		}
		return nil, serror.ErrorNotFound{Err: fmt.Errorf(err)}
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2xx response: %v", resp)
	}

	a := &agent.Agent{}

	if err := json.NewDecoder(resp.Body).Decode(a); err != nil {
		return nil, err
	}

	return a, nil
}