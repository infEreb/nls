package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/infEreb/nls/netflow/server/agent/pkg/agent"
	"github.com/infEreb/nls/netflow/server/pkg/serror"
	"golang.org/x/exp/maps"
)

type memory struct {
	sync.RWMutex
	agents map[string]*agent.Agent
	ids uint64
}

func NewRepositoryMemory() *memory {
	return &memory{
		agents: map[string]*agent.Agent{},
		ids: 0,
	}
}

func (m *memory) CreateOne(ctx context.Context, ag *agent.Agent) (*agent.Agent, error) {
	m.Lock()
	defer m.Unlock()

	if ag.ID == "" {
		ag.ID = fmt.Sprint(m.ids)
	}
	if _, has := m.agents[ag.ID]; has {
		return nil, serror.ErrorAlreadyExists{Err: fmt.Errorf("object with ID %s already exists", ag.ID)}
	}

	m.agents[ag.ID] = ag
	m.ids++
	return ag, nil
}

func (m *memory) FindOne(ctx context.Context, filters interface{}) (*agent.Agent, error) {
	m.Lock()
	defer m.Unlock()

	switch f := filters.(type) {
	case string:
		// like an ID if string type
		e, has := m.agents[f]
		if !has {
			return nil, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		return e, nil
	default:
		return nil, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}

func (m *memory) FindMany(ctx context.Context, filters interface{}) (foundAgs []*agent.Agent, err error) {
	m.Lock()
	defer m.Unlock()

	sErr := &serror.ErrorMessage{}

	if filters == nil {
		return maps.Values(m.agents), nil
	}

	switch f := filters.(type) {
	case []string:
		for _, fid := range f {
			fa, err := m.FindOne(ctx, fid)
			if err != nil {
				sErr.Errs = append(sErr.Errs, err)
			} else {
				foundAgs = append(foundAgs, fa)
			}
		}
		return foundAgs, sErr
	default:
		sErr.Errs = append(sErr.Errs, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)})
		return nil, sErr
	}
}

func (m *memory) UpdateOne(ctx context.Context, ag *agent.Agent) (*agent.Agent, error) {
	m.Lock()
	defer m.Unlock()

	if ag.ID == "" {
		return nil, serror.ErrorUndefinedField{Err: fmt.Errorf("field value is undefined: ID")}
	}

	if _, has := m.agents[ag.ID]; !has {
		return nil, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", ag.ID)}
	}

	m.agents[ag.ID] = ag

	return m.agents[ag.ID], nil
}

func (m *memory) DeleteOne(ctx context.Context, filters interface{}) (deleted uint32, err error) {
	m.Lock()
	defer m.Unlock()

	if filters == nil {
		return 0, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", filters)}
	}

	switch f := filters.(type) {
	case string:
		// like an ID if string type
		_, has := m.agents[f]
		if !has {
			return 0, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		delete(m.agents, f)
		return 1, nil
	default:
		return 0, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}