package graph

import "github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/11-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
