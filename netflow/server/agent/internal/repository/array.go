package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/infEreb/nls/netflow/server/agent/pkg/agent"
	"golang.org/x/exp/maps"
)

type array struct {
	sync.RWMutex
	agents map[string]*agent.Agent
	ids uint64
}

func NewArrayRepository() *array {
	return &array{
		agents: map[string]*agent.Agent{},
		ids: 0,
	}
}

func (a *array) CreateOne(ctx context.Context, ag *agent.Agent) (*agent.Agent, error) {
	a.Lock()
	defer a.Unlock()

	if ag.ID == "" {
		ag.ID = fmt.Sprint(a.ids)
	}
	if _, has := a.agents[ag.ID]; has {
		return nil, agent.ErrorAlreadyExists{Err: fmt.Errorf("object with ID %s already exists", ag.ID)}
	}

	a.agents[ag.ID] = ag
	a.ids++
	return ag, nil
}

func (a *array) FindOne(ctx context.Context, filters interface{}) (*agent.Agent, error) {
	a.Lock()
	defer a.Unlock()

	switch f := filters.(type) {
	case string:
		// like an ID if string type
		e, has := a.agents[f]
		if !has {
			return nil, agent.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		return e, nil
	default:
		return nil, agent.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}

func (a *array) FindMany(ctx context.Context, filters interface{}) (foundAgs []*agent.Agent, errs error) {
	a.Lock()
	defer a.Unlock()

	if filters == nil {
		return maps.Values(a.agents), nil
	}

	switch f := filters.(type) {
	case []string:
		for _, fid := range f {
			fa, err := a.FindOne(ctx, fid)
			if err != nil {
				errs = append(errs, err)
			} else {
				foundAgs = append(foundAgs, fa)
			}
		}
		return foundAgs, errs
	default:
		return nil, []error{agent.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}}
	}
}

func (a *array) UpdateOne(ctx context.Context, ag *agent.Agent) (*agent.Agent, error) {
	a.Lock()
	defer a.Unlock()

	if ag.ID == "" {
		return nil, agent.ErrorUndefinedField{Err: fmt.Errorf("field value is undefined: ID")}
	}

	if _, has := a.agents[ag.ID]; !has {
		return nil, agent.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", ag.ID)}
	}

	a.agents[ag.ID] = ag

	return a.agents[ag.ID], nil
}

func (a *array) DeleteOne(ctx context.Context, filters interface{}) (deleted uint, err error) {
	a.Lock()
	defer a.Unlock()

	if filters == nil {
		return 0, agent.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", filters)}
	}

	switch f := filters.(type) {
	case string:
		// like an ID if string type
		_, has := a.agents[f]
		if !has {
			return 0, agent.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		delete(a.agents, f)
		return 1, nil
	default:
		return 0, agent.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}