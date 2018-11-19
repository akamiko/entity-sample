package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/akamiko/entity-sample2/driver"
	ph "github.com/akamiko/entity-sample2/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connection, err := driver.ConnectSQL(dbHost, dbUser, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewUserHandler(connection)
	//r.Get("/posts", pHandler.Fetch)
	r.Get("/posts/{id}", pHandler.GetByID)
	//r.Post("/posts/create", pHandler.Create)
	//r.Put("/posts/update/{id}", pHandler.Update)
	//r.Delete("/posts/{id}", pHandler.Delete)

	fmt.Println("Server listen at :8080")
	http.ListenAndServe(":8080", r)
}
