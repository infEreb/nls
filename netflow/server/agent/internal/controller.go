package internal

import (
	"context"

	"github.com/infEreb/nls/netflow/server/agent/pkg/agent"
)

type Controller struct {
	repo Repository
}

func NewController(repo Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) CreateOne(ctx context.Context, a *agent.Agent) (*agent.Agent, error) {
	return c.repo.CreateOne(ctx, a)
}

func (c *Controller) FindByID(ctx context.Context, id string) (*agent.Agent, error) {
	return c.repo.FindOne(ctx, id)
}

func (c *Controller) FindAll(ctx context.Context) ([]*agent.Agent, []error) {
	return c.repo.FindMany(ctx, nil)
}

func (c *Controller) UpdateOne(ctx context.Context, a *agent.Agent) (*agent.Agent, error) {
	return c.repo.UpdateOne(ctx, a)
}

func (c *Controller) DeleteOne(ctx context.Context, filters interface{}) (uint, error) {
	return c.repo.DeleteOne(ctx, filters)
}
