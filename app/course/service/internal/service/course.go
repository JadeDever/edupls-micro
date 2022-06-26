package service

import (
	"context"

	pb "edupls/api/course/service/v1"
)

type CourseService struct {
	pb.UnimplementedCourseServer
}

func NewCourseService() *CourseService {
	return &CourseService{}
}

func (s *CourseService) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseReply, error) {
	return &pb.CreateCourseReply{}, nil
}
func (s *CourseService) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.UpdateCourseReply, error) {
	return &pb.UpdateCourseReply{}, nil
}
func (s *CourseService) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseReply, error) {
	return &pb.DeleteCourseReply{}, nil
}
func (s *CourseService) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseReply, error) {
	return &pb.GetCourseReply{}, nil
}
func (s *CourseService) ListCourse(ctx context.Context, req *pb.ListCourseRequest) (*pb.ListCourseReply, error) {
	return &pb.ListCourseReply{}, nil
}
