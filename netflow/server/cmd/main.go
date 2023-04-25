package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	agentR "github.com/infEreb/nls/netflow/server/app/agent/pkg/router"
	"github.com/infEreb/nls/netflow/server/app/ethernet/pkg/gateway"
	ethernetR "github.com/infEreb/nls/netflow/server/app/ethernet/pkg/router"
	packetR "github.com/infEreb/nls/netflow/server/app/packet/pkg/router"
	"github.com/infEreb/nls/netflow/server/pkg/discovery"
	"github.com/infEreb/nls/netflow/server/pkg/netflow"
	"golang.org/x/sync/errgroup"
)

const (
	AgentName = "agent"
	EthernetName = "ethernet"
	PacketName = "packet"

	AgentAddr = "localhost"
	AgentPort = "9000"

	EtherAddr = "localhost"
	EtherPort = "9001"

	PacketAddr = "localhost"
	PacketPort = "9002"
)

var (
	AgentID string
	EthernetID string
	PacketID string
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

	ethernetR.RouterGinInit(root, nil, gateway.NewAgentGatewayHTTP(registy))

	return r
}

func packetMicro(ctx context.Context, registy discovery.Registry) http.Handler {
	r := gin.Default()
	root := r.Group(netflow.ROOT)

	PacketID = discovery.GenerateInstanceID(PacketName)
	if err := registy.Register(ctx, PacketID, PacketName, PacketAddr, PacketPort); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := registy.ReportHealthyState(PacketID, PacketName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	packetR.RouterGinInit(root, nil)

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
	packetServer := http.Server{
		Addr: PacketAddr+":"+PacketPort,
		Handler: packetMicro(ctx, reg),
	}

	g.Go(func() error {
		defer reg.Deregister(ctx, AgentID, AgentName)
		return agentServer.ListenAndServe()
	})

	g.Go(func() error {
		defer reg.Deregister(ctx, EthernetID, EthernetName)
		return ethernetServer.ListenAndServe()
	})
	
	g.Go(func() error {
		defer reg.Deregister(ctx, PacketID, PacketName)
		return packetServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}