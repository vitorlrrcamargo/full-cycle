package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/15-sqlc/02-starting-with-sqlc/internal/db"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CoursesParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRB := tx.Rollback(); errRB != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRB, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CoursesParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

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

	// 	courseArgs := CoursesParams{
	// 		ID:          uuid.New().String(),
	// 		Name:        "GO",
	// 		Description: sql.NullString{String: "Go programming language", Valid: true},
	// 		Price:       100.95,
	// 	}

	// 	CategoryParams := CategoryParams{
	// 		ID:          uuid.New().String(),
	// 		Name:        "Backend",
	// 		Description: sql.NullString{String: "Backend course", Valid: true},
	// 	}

	// 	courseDB := NewCourseDB(dbConn)
	// 	err = courseDB.CreateCourseAndCategory(ctx, CategoryParams, courseArgs)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f\n",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

}
