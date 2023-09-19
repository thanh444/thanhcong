package animal

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/ent/animal"
	"golang-boilerplate/model"
	"strings"

	"github.com/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, input model.CreateAnimalInput) (*ent.Animal, error)
	List(ctx context.Context, filter model.AnimalFilterInput) (*ent.AnimalConnection, error)
}

type impl struct {
	client *ent.Client
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}

// Create implements repository.
func (a impl) Create(ctx context.Context, input model.CreateAnimalInput) (*ent.Animal, error) {
	animal, err := a.client.Animal.Create().SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return animal, nil
}

// List implements repository.
func (a impl) List(ctx context.Context, filter model.AnimalFilterInput) (*ent.AnimalConnection, error) {
	query := a.client.Animal.Query()
	if filter.Name != nil {
		query = query.Where(animal.NameContainsFold(strings.TrimSpace(*filter.Name)))
	}
	animals, err := query.Paginate(ctx, filter.Pagination.After, filter.Pagination.First, filter.Pagination.Before, filter.Pagination.Last)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return animals, nil
}
