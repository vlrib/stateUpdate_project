package main

import (
	"context"
	api "stateUpdate_project/stateUpdate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) GetProject(ctx context.Context, in *api.SearchRequest) (*api.SearchResponse, error) {
	if len(in.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Project name cannot be empty")
	}

	project, err := s.Store.GetProject(in.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get project: %v", err)
	}

	// hasFind, err := s.Store.GetProject(in.Name)
	// if err != nil {
	// 	return nil, status.Errorf(codes.InvalidArgument, "Failed to search project: %v", err)
	// }

	// if res, _ := hasFind.RowsAffected(); res != 0 {
	// 	return &api.SearchResponse{Find: 1, Message: "Project found"}, nil
	// }
	// return &api.SearchResponse{Find: 0, Message: "Project found"}, nil

	return &api.SearchResponse{Id: project.ID, Name: project.Name, Status: project.Status}, nil
}

func (s Service) CreateProject(ctx context.Context, in *api.ProjectRequest) (*api.ProjectResponse, error) {
	if len(in.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Project name cannot be empty")
	}

	if err := s.Store.CreateProject(in.Name, in.Status); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create project: %v", err)
	}
	return &api.ProjectResponse{Success: true, Message: "Project created!"}, nil

}
