package internal

import (
	"context"

	"github.com/infEreb/nls/netflow/server/app/packet/pkg/packet"
)

type Repository interface {
	CreateOne(context.Context, *packet.Packet) (*packet.Packet, error)
	FindOne(context.Context, interface{}) (*packet.Packet, error)
	FindMany(context.Context, interface{}) ([]*packet.Packet, error)
	UpdateOne(context.Context, *packet.Packet) (*packet.Packet, error)
	DeleteOne(context.Context, interface{}) (uint32, error)
}