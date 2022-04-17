package main

import (
	"fmt"
	"net/http"
	"os"
	"service-notif/db"
	"service-notif/user"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file specified")
	}
}

func main() {
	httpRouter := chi.NewRouter()
	httpRouter.Use(middleware.Logger)
	db := db.InitDB()
	httpRouter.Use(render.SetContentType(render.ContentTypeJSON))
	httpRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	httpRouter.Mount("/user", user.NewUserInjection(db).GetRoutes())

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	fmt.Println("Server running on PORT", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), httpRouter)
}
