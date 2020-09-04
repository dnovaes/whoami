package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	db "github.com/dnovaes/portfolio/database"
	"github.com/dnovaes/portfolio/gqlgen/graph/generated"
	"github.com/dnovaes/portfolio/gqlgen/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateContact(ctx context.Context, input model.NewContact) (*model.Contact, error) {
	currentTime := time.Now().UTC()
	newObjectId := primitive.NewObjectID()

	newContactDoc := &bson.D{
		{"_id", newObjectId},
		{"name", input.Name},
		{"email", input.Email},
		{"message", input.Message},
		{"createdAt", &currentTime},
	}

	collection := db.GetCollection("contacts")
	result := db.Insert(collection, *newContactDoc)
	if result.InsertedID == nil {
		return nil, errors.New("Couldn't create new contact. Please contact the developer team")
	}
	newContact := &model.Contact{
		ID:        result.InsertedID.(primitive.ObjectID),
		Name:      input.Name,
		Email:     input.Email,
		Message:   input.Message,
		CreatedAt: &currentTime,
	}
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
	return nil, nil
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id primitive.ObjectID) (*model.Contact, error) {
	deletedContact := db.FindAndDeleteContact(id)
	if deletedContact == nil {
		errorMessage := fmt.Sprintf("Couldn't delete contact: '%s' doesn't exist anymore", id.Hex())
		return deletedContact, errors.New(errorMessage)
	}
	return deletedContact, nil
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	result := db.FindAllContacts()
	return result, nil
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
