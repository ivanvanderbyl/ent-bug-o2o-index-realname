package main

import (
	"context"
	"log"
	"os"

	"github.com/ivanvanderbyl/entbug/ent"
	"github.com/ivanvanderbyl/entbug/ent/migrate"
	_ "github.com/lib/pq"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatalln("DATABASE_URL is not set")
	}

	ctx := context.Background()

	// db, _ := sql.Open("postgres", databaseURL)
	// driver := entsql.OpenDB("postgres", db)
	client, err := ent.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalln(err)
	}
	// defer db.Close()

	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalln(err)
	}
}
