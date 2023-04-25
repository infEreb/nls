package internal

import (
	"context"

	"github.com/infEreb/nls/netflow/server/app/agent/pkg/agent"
)

type Repository interface {
	CreateOne(context.Context, *agent.Agent) (*agent.Agent, error)
	FindOne(context.Context, interface{}) (*agent.Agent, error)
	FindMany(context.Context, interface{}) ([]*agent.Agent, error)
	UpdateOne(context.Context, *agent.Agent) (*agent.Agent, error)
	DeleteOne(context.Context, interface{}) (uint32, error)
}