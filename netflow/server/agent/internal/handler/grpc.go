package handler

import (
	"context"
	"errors"

	"github.com/infEreb/nls/netflow/server/agent/internal"
	"github.com/infEreb/nls/netflow/server/agent/pkg/agent"
	"github.com/infEreb/nls/netflow/server/agent/src/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HandlerGRPC struct {
	gen.AgentServiceServer
	ctrl *internal.Controller
}

func NewHandlerGRPC(ctrl *internal.Controller) *HandlerGRPC {
	return &HandlerGRPC{
		ctrl: ctrl,
	}
}

func (h *HandlerGRPC) GetAgent(ctx context.Context, req *gen.GetAgentRequest) (*gen.GetAgentResponse, error) {
	if req == nil || req.AgentId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	a, err := h.ctrl.FindByID(ctx, req.AgentId)
	if err != nil && errors.As(err, &agent.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.GetAgentResponse{Agent: agent.AgentToProto(a)}, nil
}

func (h *HandlerGRPC) GetAgents(ctx context.Context, req *gen.GetAgentsRequest) (*gen.GetAgentsResponse, error) {
	if req == nil  {
		return nil, status.Errorf(codes.InvalidArgument, "nil req")
	}
	as, err := h.ctrl.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.GetAgentResponse{Agent: agent.AgentToProto(a)}, nil
}