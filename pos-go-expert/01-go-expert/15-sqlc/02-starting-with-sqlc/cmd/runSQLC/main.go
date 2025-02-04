package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/15-sqlc/02-starting-with-sqlc/internal/db"
)

func main() {
	// Open a connection to the database
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	// Create a new db instance
	queries := db.New(dbConn)

	// Create a new category
	id := uuid.New().String()

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          id,
		Name:        "Backend",
		Description: sql.NullString{String: "Backend Description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	PrintCategories(*queries, ctx)

	// Update a category
	queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          id,
		Name:        "Backend Updated",
		Description: sql.NullString{String: "Backend Description Updated", Valid: true},
	})

	PrintCategories(*queries, ctx)

	// Delete a category
	err = queries.DeleteCategory(ctx, id)

	if err != nil {
		panic(err)
	}

	PrintCategories(*queries, ctx)
}

func PrintCategories(queries db.Queries, ctx context.Context) {
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println("id:", category.ID, "| name:", category.Name, "| description:", category.Description.String)
	}
}
