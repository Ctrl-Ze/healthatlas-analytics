package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Ctrl-Ze/healthatlas-analytics/internal/consumer"
	"github.com/Ctrl-Ze/healthatlas-analytics/internal/db"
)

func main() {
	ctx := context.Background()

	dbURL := db.BuildPostgresURL()

	if err := db.RunMigrations(dbURL); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	pool, err := db.NewPool(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer pool.Close()

	if os.Getenv("DISABLE_KAFKA") != "true" {
		repo := db.NewRepository(pool)
		c := consumer.New(
			[]string{os.Getenv("KAFKA_BROKER")},
			"blood-test.recorded",
			"themis-group",
			repo,
		)
		c.Start(ctx)

	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, e *http.Request) {
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, e *http.Request) {
		w.Write([]byte("Hello from Themis!"))
	})

	log.Println("Themis listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
