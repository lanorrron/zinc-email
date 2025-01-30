package main

import (
	"awesomeProject/config"
	"awesomeProject/internal/email/handler"
	"awesomeProject/internal/email/repository"
	"awesomeProject/internal/email/routes"
	"awesomeProject/internal/email/service"
	"awesomeProject/internal/zincsearch"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.Use(middleware.Logger)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.LoadConfig()

	client := zincsearch.NewZincClient(cfg.ZincSearchHost, cfg.ZincSearchUser, cfg.ZincSearchPassword)
	emailRepo := repository.NewEmailRepository(client)
	emailService := service.NewEmailService(emailRepo)
	emailHandler := handler.NewEmailHandler(emailService)

	routes.InitializeMailRoutes(r, emailHandler)

	fmt.Println("Listening on port", cfg.ServerPort)
	http.ListenAndServe(cfg.ServerPort, r)
}
