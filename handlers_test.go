package main

import (
	"context"
	"errors"
	api "stateUpdate_project/stateUpdate"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FakeStore struct {
	Error   error
	project Project
}

func (s *FakeStore) CreateProject(name string, stat string) error {
	return s.Error
}

func (s *FakeStore) GetProject(name string) (project Project, err error) {
	return s.project, s.Error
}

//CreateProject Test
func TestCreatProjectEmptyName(t *testing.T) {
	svc := Service{}
	in := &api.ProjectRequest{}

	_, err := svc.CreateProject(context.Background(), in)
	assert.Equal(t, status.Code(err), codes.InvalidArgument)
}

func TestCreatProjectInternalError(t *testing.T) {
	svc := Service{Store: &FakeStore{Error: errors.New("error")}}
	in := &api.ProjectRequest{Name: "teste123"}

	_, err := svc.CreateProject(context.Background(), in)
	assert.Equal(t, status.Code(err), codes.Internal)
}
func TestCreateProjectReturn(t *testing.T) {
	svc := Service{Store: &FakeStore{}}
	in := &api.ProjectRequest{Name: "teste123"}

	resp, err := svc.CreateProject(context.Background(), in)
	assert.Nil(t, err)
	assert.True(t, resp.Success)
}

//CreateProject Test

//SearchProject Test
func TestGetProjectEmptyName(t *testing.T) {
	svc := Service{Store: &FakeStore{}}
	in := &api.SearchRequest{Name: ""}

	_, err := svc.GetProject(context.Background(), in)
	assert.Equal(t, status.Code(err), codes.Internal)
}

func TestGetProjectOk(t *testing.T) {
	svc := Service{Store: &FakeStore{project: Project{Name: "test"}}}
	in := &api.SearchRequest{Name: "test"}

	res, _ := svc.GetProject(context.Background(), in)
	assert.Equal(t, res.Name, "test")
}

func TestGetProjectInternalError(t *testing.T) {
	svc := Service{Store: &FakeStore{Error: errors.New("failed")}}
	in := &api.SearchRequest{Name: "test"}

	_, err := svc.GetProject(context.Background(), in)
	if err != nil {

	}
}

func TestGetProjectReturn(t *testing.T) {
	svc := Service{Store: &FakeStore{project: Project{ID: "123", Name: "teste", Status: "1"}}}
	in := &api.SearchRequest{Name: "test"}

	res, err := svc.GetProject(context.Background(), in)
	assert.Equal(t, res.Id, "123")
	assert.Equal(t, res.Name, "teste")
	assert.Equal(t, res.Status, "1")
	assert.Nil(t, err)
}
