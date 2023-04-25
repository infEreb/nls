package handler

import (
	"context"
	"errors"

	"github.com/infEreb/nls/netflow/server/app/ethernet/internal"
	"github.com/infEreb/nls/netflow/server/app/ethernet/pkg/ethernet"
	"github.com/infEreb/nls/netflow/server/app/ethernet/src/gen"
	"github.com/infEreb/nls/netflow/server/pkg/serror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HandlerGRPC struct {
	gen.EthernetServiceServer
	ctrl *internal.Controller
}

func NewHandlerGRPC(ctrl *internal.Controller) *HandlerGRPC {
	return &HandlerGRPC{
		ctrl: ctrl,
	}
}

func (h *HandlerGRPC) GetEthernet(ctx context.Context, req *gen.GetEthernetRequest) (*gen.GetEthernetResponse, error) {
	if req == nil || req.EthernetId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or empty id")
	}
	a, err := h.ctrl.FindByID(ctx, req.EthernetId)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.GetEthernetResponse{Ethernet: ethernet.EthernetToProto(a)}, nil
}

func (h *HandlerGRPC) GetEthernets(ctx context.Context, req *gen.GetEthernetsRequest) (*gen.GetEthernetsResponse, error) {
	if req == nil  {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}
	as, err := h.ctrl.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pas := []*gen.Ethernet{}
	for _, a := range as {
		pas = append(pas, ethernet.EthernetToProto(a))
	}

	return &gen.GetEthernetsResponse{Ethernets: pas}, nil
}

func (h *HandlerGRPC) PostEthernet(ctx context.Context, req *gen.PostEthernetRequest) (*gen.PostEthernetResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	a, err := h.ctrl.CreateOne(ctx, ethernet.EthernetFromProto(req.Ethernet))
	if err != nil && errors.As(err, &serror.ErrorAlreadyExists{}) {
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.PostEthernetResponse{Ethernet: ethernet.EthernetToProto(a)}, nil
}

func (h *HandlerGRPC) PutEthernet(ctx context.Context, req *gen.PutEthernetRequest) (*gen.PutEthernetResponse, error) {
	if req == nil || req.EthernetId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or id")
	}

	a, err := h.ctrl.UpdateOne(ctx, ethernet.EthernetFromProto(req.Ethernet))
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil && errors.As(err, &serror.ErrorUndefinedField{}) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.PutEthernetResponse{Ethernet: ethernet.EthernetToProto(a)}, nil
}

func (h *HandlerGRPC) DeleteEthernet(ctx context.Context, req *gen.DeleteEthernetRequest) (*gen.DeleteEthernetResponse, error) {
	if req == nil || req.EthernetId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or id")
	}

	n, err := h.ctrl.DeleteOne(ctx, req.EthernetId)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil && errors.As(err, &serror.ErrorUnexpectedType{}) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.DeleteEthernetResponse{Deleted: n}, nil
} 