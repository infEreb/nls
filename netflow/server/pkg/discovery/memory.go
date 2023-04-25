package discovery

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/infEreb/nls/netflow/server/pkg/serror"
)

type RegistryMemory struct {
	sync.RWMutex
	// map[serviceName][instanceID]
	serviceAddrs map[string]map[string]*serviceInstance
}

type serviceInstance struct {
	address string
	port string
	lastActive time.Time
}

func NewRegistryMemory() *RegistryMemory {
	return &RegistryMemory{
		serviceAddrs: map[string]map[string]*serviceInstance{},
	}
}

func (r *RegistryMemory) Register(ctx context.Context, instanceID string, serviceName string, addr string, port string) error {
	r.Lock()
	defer r.Unlock()

	if _, has := r.serviceAddrs[serviceName]; !has {
		r.serviceAddrs[serviceName] = map[string]*serviceInstance{}
	}
	r.serviceAddrs[serviceName][instanceID] = &serviceInstance{
		address: addr,
		port: port,
		lastActive: time.Now(),
	}

	return nil
}

func (r *RegistryMemory) Deregister(ctx context.Context, instanceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, has := r.serviceAddrs[serviceName]; !has {
		return serror.ErrorNotFound{Err: fmt.Errorf("not found service %s", serviceName)}
	}

	delete(r.serviceAddrs[serviceName], instanceID)
	return nil
}

func (r *RegistryMemory) ReportHealthyState(instanceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, has := r.serviceAddrs[serviceName]; !has {
		return serror.ErrorNotFound{Err: fmt.Errorf("not found service %s", serviceName)}
	}
	if _, has := r.serviceAddrs[serviceName][instanceID]; !has {
		return serror.ErrorNotFound{Err: fmt.Errorf("not found instance %s of service %s", instanceID, serviceName)}
	}

	r.serviceAddrs[serviceName][instanceID].lastActive = time.Now()
	return nil
}

func (r *RegistryMemory) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	if len(r.serviceAddrs[serviceName]) == 0 {
		return nil, serror.ErrorNotFound{Err: fmt.Errorf("not found service %s", serviceName)}
	}

	res := []string{}

	for _, i := range r.serviceAddrs[serviceName] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}
		res = append(res, fmt.Sprintf("%s:%s", i.address, i.port))
	}
	return res, nil
}