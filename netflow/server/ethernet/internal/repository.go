package internal

import (
	"context"

	"github.com/infEreb/nls/netflow/server/ethernet/pkg/ethernet"
)

type Repository interface {
	CreateOne(context.Context, *ethernet.Ethernet) (*ethernet.Ethernet, error)
	CreateMany(context.Context, []*ethernet.Ethernet) ([]*ethernet.Ethernet, []error)
	FindOne(context.Context, interface{}) (*ethernet.Ethernet, error)
	FindMany(context.Context, interface{}) ([]*ethernet.Ethernet, []error)
}