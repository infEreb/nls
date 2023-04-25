package gateway

import (
	"context"

	"github.com/infEreb/nls/netflow/server/app/agent/pkg/agent"
)

type GatewayAgent interface {
	GetAgent(context.Context, string) (*agent.Agent, error)
}