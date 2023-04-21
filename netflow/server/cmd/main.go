package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	agentR "github.com/infEreb/nls/netflow/server/agent/pkg/router"
	"github.com/infEreb/nls/netflow/server/ethernet/pkg/gateway"
	ethernetR "github.com/infEreb/nls/netflow/server/ethernet/pkg/router"
	"github.com/infEreb/nls/netflow/server/pkg/discovery"
	"github.com/infEreb/nls/netflow/server/pkg/netflow"
	"golang.org/x/sync/errgroup"
)

const (
	AgentName = "agent"
	EthernetName = "ethernet"

	AgentAddr = "localhost"
	AgentPort = "9000"

	EtherAddr = "localhost"
	EtherPort = "9001"
)

var (
	AgentID string
	EthernetID string
)

func agentMicro(ctx context.Context, registy discovery.Registry) http.Handler {
	r := gin.Default()
	root := r.Group(netflow.ROOT)
	agentR.RouterGinInit(root, nil)

	AgentID = discovery.GenerateInstanceID(AgentName)
	if err := registy.Register(ctx, AgentID, AgentName, AgentAddr, AgentPort); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := registy.ReportHealthyState(AgentID, AgentName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return r
}

func ethernetMicro(ctx context.Context, registy discovery.Registry) http.Handler {
	r := gin.Default()
	root := r.Group(netflow.ROOT)

	EthernetID = discovery.GenerateInstanceID(EthernetName)
	if err := registy.Register(ctx, EthernetID, EthernetName, EtherAddr, EtherPort); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := registy.ReportHealthyState(EthernetID, EthernetName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	ethernetR.RouterGinInit(root, nil, gateway.NewHTTP(registy))

	return r
}

var g errgroup.Group

func main() {
	reg := discovery.NewRegistryMemory()

	ctx := context.Background()

	agentServer := http.Server{
		Addr: AgentAddr+":"+AgentPort,
		Handler: agentMicro(ctx, reg),
	}
	ethernetServer := http.Server{
		Addr: EtherAddr+":"+EtherPort,
		Handler: ethernetMicro(ctx, reg),
	}

	g.Go(func() error {
		defer reg.Deregister(ctx, AgentID, AgentName)
		return agentServer.ListenAndServe()
	})

	g.Go(func() error {
		defer reg.Deregister(ctx, EthernetID, EthernetName)
		return ethernetServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}