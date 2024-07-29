package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	infradb "go.mod/connect"
)

func main() {
	fmt.Println("Starting")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	r := chi.NewRouter()
	m := chiprometheus.NewMiddleware("router")
	infradb.DSN = "postgresql://user:password@postgres:5451/database"

	r.Use(m)
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	r.Use(corsOptions.Handler)
	r.Mount("/user", UserRouter())
	r.Mount("/campaign", CampaignRouter())

	port := envPortOr("8080")
	fmt.Println(":", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		fmt.Println(envPort)
		return ":" + envPort
	}
	return ":" + port
}
