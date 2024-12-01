package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Joshdike/expense-tracker/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	connString := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil{
		log.Fatal(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	h := handlers.New(pool)

	r.Get("/expense", h.GetAllExpense)
	r.Get("/expense/{id}", h.GetExpenseById)
	r.Post("/expense", h.CreateExpense)
	r.Put("/expense/{id}", h.UpdateExpense)
	r.Delete("/expense/{id}", h.DeleteExpense)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}

}
