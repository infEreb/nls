package internal

import (
	"context"

	"github.com/infEreb/nls/netflow/server/app/agent/pkg/agent"
	"github.com/infEreb/nls/netflow/server/app/ethernet/internal/gateway"
	"github.com/infEreb/nls/netflow/server/app/ethernet/pkg/ethernet"
)

type Controller struct {
	repo Repository
	agent gateway.GatewayAgent
}

func NewController(repo Repository, agentGateway gateway.GatewayAgent) *Controller {
	return &Controller{
		repo: repo,
		agent: agentGateway,
	}
}

func (c *Controller) CreateOne(ctx context.Context, e *ethernet.Ethernet) (*ethernet.Ethernet, error) {
	return c.repo.CreateOne(ctx, e)
}

func (c *Controller) FindByID(ctx context.Context, id string) (*ethernet.Ethernet, error) {
	return c.repo.FindOne(ctx, id)
}
func (c *Controller) FindAll(ctx context.Context) ([]*ethernet.Ethernet, error) {
	return c.repo.FindMany(ctx, nil)
}

func (c *Controller) UpdateOne(ctx context.Context, e *ethernet.Ethernet) (*ethernet.Ethernet, error) {
	return c.repo.UpdateOne(ctx, e)
}

func (c *Controller) DeleteOne(ctx context.Context, id string) (uint32, error) {
	return c.repo.DeleteOne(ctx, id)
}

func (c *Controller) GetAgent(ctx context.Context, agentID string) (*agent.Agent, error) {
	return c.agent.GetAgent(ctx, agentID)
}