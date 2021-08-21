package persist

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func Initialize() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")

	if err != nil {
		panic(fmt.Sprintf("%s, Error: %s", "Failed to connect.", err))
	}
	defer conn.Close(context.Background())

	sql := "select id, name from Test"
	var id, name string
	scanErr := conn.QueryRow(context.Background(), sql).Scan(&id, &name)

	if scanErr != nil {
		panic(fmt.Sprintf("%s, Error: %s", "Failed to execute query.", scanErr))

	}

	fmt.Println("Found data...")
	fmt.Printf("ID: %s, Name: %s", id, name)

}
