package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/ksuid"
)

type Store interface {
	CreateProject(name string, stat string) error
	GetProject(name string) (project Project, err error)
}

type StoreImpl struct {
	DB *sqlx.DB
}

func (s StoreImpl) GetProject(name string) (project Project, err error) {
	if err = s.DB.Get(&project, SelectProjectByName, name); err != nil {
		return Project{}, err
	}

	return project, nil
}

func (s StoreImpl) CreateProject(name string, stat string) error {
	_, err := s.DB.NamedExec(CreateProject, &Project{
		ID:     ksuid.New().String(),
		Name:   name,
		Status: stat,
	})

	return err
}
