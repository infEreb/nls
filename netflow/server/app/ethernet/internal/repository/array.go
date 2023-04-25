package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/infEreb/nls/netflow/server/app/ethernet/pkg/ethernet"
	"github.com/infEreb/nls/netflow/server/pkg/serror"
	"golang.org/x/exp/maps"
)

type memory struct {
	sync.RWMutex
	eths map[string]*ethernet.Ethernet
}

func NewRepositoryMemory() *memory {
	return &memory{
		eths: map[string]*ethernet.Ethernet{},
	}
}

func (m *memory) CreateOne(ctx context.Context, eth *ethernet.Ethernet) (*ethernet.Ethernet, error) {
	m.Lock()
	defer m.Unlock()

	if eth.ID == "" {
		eth.ID = uuid.NewString()
	}
	if _, has := m.eths[eth.ID]; has {
		return nil, serror.ErrorAlreadyExists{Err: fmt.Errorf("object with ID %s already exists", eth.ID)}
	}

	m.eths[eth.ID] = eth
	return eth, nil
}
func (m *memory) CreateMany(ctx context.Context, eths []*ethernet.Ethernet) (createdEths []*ethernet.Ethernet, err error) {
	m.Lock()
	defer m.Unlock()

	errs := &serror.ErrorMessage{}

	for _, e := range eths {
		ce, err := m.CreateOne(ctx, e)
		if err != nil {
			errs.Errs = append(errs.Errs, err)
		} else {
			createdEths = append(createdEths, ce)
		}
	}

	return createdEths, errs
}
func (m *memory) FindOne(ctx context.Context, filters interface{}) (*ethernet.Ethernet, error) {
	m.Lock()
	defer m.Unlock()

	switch f := filters.(type) {
	case string:
		// like an ID if string type
		e, has := m.eths[f]
		if !has {
			return nil, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		return e, nil
	default:
		return nil, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}
func (m *memory) FindMany(ctx context.Context, filters interface{}) (foundEths []*ethernet.Ethernet, err error) {
	m.Lock()
	defer m.Unlock()

	if filters == nil {
		return maps.Values(m.eths), nil
	}

	errs := &serror.ErrorMessage{}

	switch f := filters.(type) {
	case []string:
		for _, fid := range f {
			fe, err := m.FindOne(ctx, fid)
			if err != nil {
				errs.Errs = append(errs.Errs, err)
			} else {
				foundEths = append(foundEths, fe)
			}
		}
		return foundEths, errs
	default:
		errs.Errs = append(errs.Errs, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)})
		return nil, errs
	}
}

func (m *memory) UpdateOne(ctx context.Context, eth *ethernet.Ethernet) (*ethernet.Ethernet, error) {
	m.Lock()
	defer m.Unlock()

	if eth.ID == "" {
		return nil, serror.ErrorUndefinedField{Err: fmt.Errorf("field value is undefined: ID")}
	}

	if _, has := m.eths[eth.ID]; !has {
		return nil, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", eth.ID)}
	}

	m.eths[eth.ID] = eth

	return m.eths[eth.ID], nil
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
		_, has := m.eths[f]
		if !has {
			return 0, serror.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		delete(m.eths, f)
		return 1, nil
	default:
		return 0, serror.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}
