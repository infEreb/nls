package gateway

import (
	"context"

	"github.com/infEreb/nls/netflow/server/agent/pkg/agent"
)

type GatewayAgent interface {
	GetAgent(context.Context, string) (*agent.Agent, error)
}