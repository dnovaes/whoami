package graph

import "github.com/dnovaes/portfolio/gqlgen/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	contacts    []*model.Contact
	experiences []*model.Experience
}
