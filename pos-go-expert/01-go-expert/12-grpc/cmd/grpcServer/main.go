package main

import (
	"database/sql"
	"net"

	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/12-grpc/internal/database"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/12-grpc/internal/pb"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/12-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS "categories" (
		"id"			TEXT,
		"name"			TEXT,
		"description"	TEXT
	)`)

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
