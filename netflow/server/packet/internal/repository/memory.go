package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/infEreb/nls/netflow/server/packet/pkg/packet"
	"github.com/infEreb/nls/netflow/server/pkg/serror"
	"golang.org/x/exp/maps"
)

type memory struct {
	sync.RWMutex
	packets map[string]*packet.Packet
	ids uint64
}

func NewRepositoryMemory() *memory {
	return &memory{
		packets: map[string]*packet.Packet{},
		ids: 0,
	}
}

func (m *memory) CreateOne(ctx context.Context, pck *packet.Packet) (*packet.Packet, error) {
	m.Lock()
	defer m.Unlock()

	if pck.ID == "" {
		pck.ID = fmt.Sprint(m.ids)
	}
	if _, has := m.packets[pck.ID]; has {
		return nil, serror.ErrorAlreadyExists{Err: fmt.Errorf("object with ID %s already exists", pck.ID)}
	}

	m.packets[pck.ID] = pck
	m.ids++
	return pck, nil
}

func (m *memory) FindOne(ctx context.Context, filters interface{}) (*packet.Packet, error) {
	m.Lock()
	defer m.Unlock()

	switch f := filters.(type) {
	case string:
		// like an ID if string type
		e, has := m.packets[f]
		if !has {
			return nil, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		return e, nil
	default:
		return nil, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}

func (m *memory) FindMany(ctx context.Context, filters interface{}) (foundAgs []*packet.Packet, err error) {
	m.Lock()
	defer m.Unlock()

	sErr := &serror.ErrorMessage{}

	if filters == nil {
		return maps.Values(m.packets), nil
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

func (m *memory) UpdateOne(ctx context.Context, ag *packet.Packet) (*packet.Packet, error) {
	m.Lock()
	defer m.Unlock()

	if ag.ID == "" {
		return nil, serror.ErrorUndefinedField{Err: fmt.Errorf("field value is undefined: ID")}
	}

	if _, has := m.packets[ag.ID]; !has {
		return nil, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", ag.ID)}
	}

	m.packets[ag.ID] = ag

	return m.packets[ag.ID], nil
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
		_, has := m.packets[f]
		if !has {
			return 0, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		delete(m.packets, f)
		return 1, nil
	default:
		return 0, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}