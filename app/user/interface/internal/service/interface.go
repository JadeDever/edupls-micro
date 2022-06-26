package service

import (
	"context"

	pb "edupls/api/interface/service/v1"
)

type InterfaceService struct {
	pb.UnimplementedInterfaceServer
}

func NewInterfaceService() *InterfaceService {
	return &InterfaceService{}
}

func (s *InterfaceService) CreateInterface(ctx context.Context, req *pb.CreateInterfaceRequest) (*pb.CreateInterfaceReply, error) {
	return &pb.CreateInterfaceReply{}, nil
}
func (s *InterfaceService) UpdateInterface(ctx context.Context, req *pb.UpdateInterfaceRequest) (*pb.UpdateInterfaceReply, error) {
	return &pb.UpdateInterfaceReply{}, nil
}
func (s *InterfaceService) DeleteInterface(ctx context.Context, req *pb.DeleteInterfaceRequest) (*pb.DeleteInterfaceReply, error) {
	return &pb.DeleteInterfaceReply{}, nil
}
func (s *InterfaceService) GetInterface(ctx context.Context, req *pb.GetInterfaceRequest) (*pb.GetInterfaceReply, error) {
	return &pb.GetInterfaceReply{}, nil
}
func (s *InterfaceService) ListInterface(ctx context.Context, req *pb.ListInterfaceRequest) (*pb.ListInterfaceReply, error) {
	return &pb.ListInterfaceReply{}, nil
}
