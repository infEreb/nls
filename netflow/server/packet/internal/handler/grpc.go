package handler

import (
	"context"
	"errors"

	"github.com/infEreb/nls/netflow/server/packet/internal"
	"github.com/infEreb/nls/netflow/server/packet/pkg/packet"
	"github.com/infEreb/nls/netflow/server/packet/src/gen"
	"github.com/infEreb/nls/netflow/server/pkg/serror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HandlerGRPC struct {
	gen.PacketServiceServer
	ctrl *internal.Controller
}

func NewHandlerGRPC(ctrl *internal.Controller) *HandlerGRPC {
	return &HandlerGRPC{
		ctrl: ctrl,
	}
}

func (h *HandlerGRPC) GetPacket(ctx context.Context, req *gen.GetPacketRequest) (*gen.GetPacketResponse, error) {
	if req == nil || req.PacketId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or empty id")
	}
	p, err := h.ctrl.FindByID(ctx, req.PacketId)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.GetPacketResponse{Packet: packet.PacketToProto(p)}, nil
}

func (h *HandlerGRPC) GetPackets(ctx context.Context, req *gen.GetPacketsRequest) (*gen.GetPacketsResponse, error) {
	if req == nil  {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}
	ps, err := h.ctrl.FindAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pas := []*gen.Packet{}
	for _, a := range ps {
		pas = append(pas, packet.PacketToProto(a))
	}

	return &gen.GetPacketsResponse{Packets: pas}, nil
}

func (h *HandlerGRPC) PostPacket(ctx context.Context, req *gen.PostPacketRequest) (*gen.PostPacketResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	p, err := h.ctrl.CreateOne(ctx, packet.PacketFromProto(req.Packet))
	if err != nil && errors.As(err, &serror.ErrorAlreadyExists{}) {
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.PostPacketResponse{Packet: packet.PacketToProto(p)}, nil
}

func (h *HandlerGRPC) PutPacket(ctx context.Context, req *gen.PutPacketRequest) (*gen.PutPacketResponse, error) {
	if req == nil || req.PacketId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or id")
	}

	p, err := h.ctrl.UpdateOne(ctx, packet.PacketFromProto(req.Packet))
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil && errors.As(err, &serror.ErrorUndefinedField{}) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.PutPacketResponse{Packet: packet.PacketToProto(p)}, nil
}

func (h *HandlerGRPC) DeletePacket(ctx context.Context, req *gen.DeletePacketRequest) (*gen.DeletePacketResponse, error) {
	if req == nil || req.PacketId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or id")
	}

	n, err := h.ctrl.DeleteOne(ctx, req.PacketId)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil && errors.As(err, &serror.ErrorUnexpectedType{}) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.DeletePacketResponse{Deleted: n}, nil
} 