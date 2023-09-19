package repository

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/repository/animal"
	"golang-boilerplate/repository/user"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	User() user.Repository
	Animal() animal.Repository
}

// impl is the implementation of the repository registry.
type impl struct {
	user   user.Repository
	animal animal.Repository
}

// Animal implements RepositoryRegistry.
func (a impl) Animal() animal.Repository {
	return a.animal
}

// New creates a new repository registry.
func New(client *ent.Client) RepositoryRegistry {
	return &impl{
		user:   user.New(client),
		animal: animal.New(client),
	}
}

// User returns the user repository.
func (i impl) User() user.Repository {
	return i.user
}
