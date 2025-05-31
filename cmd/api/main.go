package main

import (
	"context"
	"fmt"
	"os"

	db "github.com/Myronarty/Lab_Go/db/sqlc"
	"github.com/Myronarty/Lab_Go/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), "postgresql://user:password@localhost:5432/koguts")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer connPool.Close()

	store := db.NewStore(connPool)

	server := server.NewServer(store)
	fmt.Println("Server is starting on port 3000...")
	server.Run(":3000")
	fmt.Println("Server stopped running")
}
