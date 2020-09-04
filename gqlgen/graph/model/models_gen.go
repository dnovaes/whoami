// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Message   string             `json:"message"`
	CreatedAt *time.Time         `json:"createdAt"`
}

type Experience struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Year        int                `json:"year"`
	Company     string             `json:"company"`
	CompanyLink *string            `json:"companyLink"`
	StartedAt   *time.Time         `json:"startedAt"`
	FinishedAt  *time.Time         `json:"finishedAt"`
}

type NewContact struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type NewExperience struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Year        int        `json:"year"`
	Company     string     `json:"company"`
	CompanyLink *string    `json:"companyLink"`
	StartedAt   *time.Time `json:"startedAt"`
	FinishedAt  *time.Time `json:"finishedAt"`
}
