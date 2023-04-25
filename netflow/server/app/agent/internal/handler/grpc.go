package handler

import (
	"context"
	"errors"

	"github.com/infEreb/nls/netflow/server/app/agent/internal"
	"github.com/infEreb/nls/netflow/server/app/agent/pkg/agent"
	"github.com/infEreb/nls/netflow/server/app/agent/src/gen"
	"github.com/infEreb/nls/netflow/server/pkg/serror"
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
		return nil, status.Errorf(codes.InvalidArgument, "nil request or empty id")
	}
	a, err := h.ctrl.FindByID(ctx, req.AgentId)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.GetAgentResponse{Agent: agent.AgentToProto(a)}, nil
}

func (h *HandlerGRPC) GetAgents(ctx context.Context, req *gen.GetAgentsRequest) (*gen.GetAgentsResponse, error) {
	if req == nil  {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}
	as, err := h.ctrl.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pas := []*gen.Agent{}
	for _, a := range as {
		pas = append(pas, agent.AgentToProto(a))
	}

	return &gen.GetAgentsResponse{Agents: pas}, nil
}

func (h *HandlerGRPC) PostAgent(ctx context.Context, req *gen.PostAgentRequest) (*gen.PostAgentResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	a, err := h.ctrl.CreateOne(ctx, agent.AgentFromProto(req.Agent))
	if err != nil && errors.As(err, &serror.ErrorAlreadyExists{}) {
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.PostAgentResponse{Agent: agent.AgentToProto(a)}, nil
}

func (h *HandlerGRPC) PutAgent(ctx context.Context, req *gen.PutAgentRequest) (*gen.PutAgentResponse, error) {
	if req == nil || req.AgentId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or id")
	}

	a, err := h.ctrl.UpdateOne(ctx, agent.AgentFromProto(req.Agent))
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil && errors.As(err, &serror.ErrorUndefinedField{}) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.PutAgentResponse{Agent: agent.AgentToProto(a)}, nil
}

func (h *HandlerGRPC) DeleteAgent(ctx context.Context, req *gen.DeleteAgentRequest) (*gen.DeleteAgentResponse, error) {
	if req == nil || req.AgentId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or id")
	}

	n, err := h.ctrl.DeleteOne(ctx, req.AgentId)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil && errors.As(err, &serror.ErrorUnexpectedType{}) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.DeleteAgentResponse{Deleted: n}, nil
} 