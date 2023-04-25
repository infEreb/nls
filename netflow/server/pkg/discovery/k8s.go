package discovery

import (
	"context"
	"fmt"
	"os"
)

type RegistryK8s struct {
	namespace string
}

func NewRegistryK8s(namespace string) *RegistryK8s {
	return &RegistryK8s{
		namespace: namespace,
	}
}

func (r *RegistryK8s) Register(ctx context.Context, serviceName string, instanceID string, addr string, port string) error {
	return nil
}

func (r *RegistryK8s) Deregister(ctx context.Context, serviceName string, instanceID string) error {
	return nil
}

func (r *RegistryK8s) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	addr, has := os.LookupEnv(fmt.Sprintf("%s_SERVICE_HOST", serviceName))
	if !has {
		return nil, fmt.Errorf("error to find %s_SERVICE_HOST env", serviceName)
	}
	port, has := os.LookupEnv(fmt.Sprintf("%s_SERVICE_PORT", serviceName))
	if !has {
		return nil, fmt.Errorf("error to find %s_SERVICE_PORT env", serviceName)
	}

	return []string{
		fmt.Sprintf("%s:%s", addr, port),
	}, nil
}

func (r *RegistryK8s) ReportHealthyState(instanceID string, serviceName string) error {
	return nil
}