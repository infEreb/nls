package gateway

import (
	"github.com/infEreb/nls/netflow/server/app/ethernet/internal/gateway"
	"github.com/infEreb/nls/netflow/server/pkg/discovery"
)

func NewAgentGatewayHTTP(registry discovery.Registry) gateway.GatewayAgent {
	return gateway.NewAgentGatewayHTTP(registry)
}