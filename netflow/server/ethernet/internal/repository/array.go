package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/infEreb/nls/netflow/server/ethernet/pkg/ethernet"
	"golang.org/x/exp/maps"
)

type array struct {
	eths map[string]*ethernet.Ethernet
}

func NewArrayRepository() *array {
	return &array{
		eths: map[string]*ethernet.Ethernet{},
	}
}

func (a *array) CreateOne(ctx context.Context, eth *ethernet.Ethernet) (*ethernet.Ethernet, error) {
	if eth.ID == "" {
		eth.ID = uuid.NewString()
	}
	if _, has := a.eths[eth.ID]; has {
		return nil, ethernet.ErrorAlreadyExists{Err: fmt.Errorf("object with ID %s already exists", eth.ID)}
	}

	a.eths[eth.ID] = eth
	return eth, nil
}
func (a *array) CreateMany(ctx context.Context, eths []*ethernet.Ethernet) (createdEths []*ethernet.Ethernet, errs []error) {
	for _, e := range eths {
		ce, err := a.CreateOne(ctx, e)
		if err != nil {
			errs = append(errs, err)
		} else {
			createdEths = append(createdEths, ce)
		}
	}

	return createdEths, errs
}
func (a *array) FindOne(ctx context.Context, filters interface{}) (*ethernet.Ethernet, error) {
	switch f := filters.(type) {
	case string:
		// like an ID if string type
		e, has := a.eths[f]
		if !has {
			return nil, ethernet.ErrorNotFound{Err: fmt.Errorf("object with ID %s not found", f)}
		}
		return e, nil
	default:
		return nil, ethernet.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}
	}
}
func (a *array) FindMany(ctx context.Context, filters interface{}) (foundEths []*ethernet.Ethernet, errs []error) {
	if filters == nil {
		return maps.Values(a.eths), nil
	}

	switch f := filters.(type) {
	case []string:
		for _, fid := range f {
			fe, err := a.FindOne(ctx, fid)
			if err != nil {
				errs = append(errs, err)
			} else {
				foundEths = append(foundEths, fe)
			}
		}
		return foundEths, errs
	default:
		return nil, []error{ethernet.ErrorUnexpectedType{Err: fmt.Errorf("unexpected filter type: %T", f)}}
	}
}
