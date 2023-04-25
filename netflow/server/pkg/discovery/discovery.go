package discovery

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Registry interface {
	Register(ctx context.Context, instanceID string, serviceName string, addr string, port string) error
	Deregister(ctx context.Context, instanceID string, serviceName string) error
	ServiceAddresses(ctx context.Context, serviceName string) ([]string, error)
	ReportHealthyState(instanceID string, serviceName string) error
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s:%s", serviceName, uuid.NewString())
}