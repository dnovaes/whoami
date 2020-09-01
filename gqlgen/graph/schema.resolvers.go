package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dnovaes/portfolio/gqlgen/graph/generated"
	"github.com/dnovaes/portfolio/gqlgen/graph/model"
)

func (r *mutationResolver) CreateContact(ctx context.Context, input model.NewContact) (*model.Contact, error) {
	for i := 0; i < len(r.contacts); i++ {
		if r.contacts[i].Message == input.Message {
			return nil, errors.New("contact Message already exists")
		}
	}

	currentTime := time.Now().UTC()
	newContact := &model.Contact{
		ID:        fmt.Sprintf("T%d", rand.Int()),
		Name:      input.Name,
		Email:     input.Email,
		Message:   input.Message,
		CreatedAt: &currentTime,
	}
	r.contacts = append(r.contacts, newContact)
	return newContact, nil
}

func (r *mutationResolver) CreateExperience(ctx context.Context, input model.NewExperience) (*model.Experience, error) {
	for i := 0; i < len(r.experiences); i++ {
		exp := r.experiences[i]
		if exp.Title == input.Title || exp.Description == input.Description {
			errorMessage := "title or description of experience already exists"
			return nil, errors.New(errorMessage)
		}
	}
	newExperience := &model.Experience{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Title:       input.Title,
		Description: input.Description,
		Year:        input.Year,
		Company:     input.Company,
		CompanyLink: input.CompanyLink,
		StartedAt:   input.StartedAt,
		FinishedAt:  input.FinishedAt,
	}
	r.experiences = append(r.experiences, newExperience)
	return newExperience, nil
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	return r.contacts, nil
}

func (r *queryResolver) Experiences(ctx context.Context) ([]*model.Experience, error) {
	return r.experiences, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
