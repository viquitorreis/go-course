package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"gitlab.com/victorreisprog/go/-/tree/master/go-free-code-camp/projetos/RSS-aggregator/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORTA não encontrada no .env") // vai parar o programa imediatamente com codigo 1 e uma mensagem
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL não encontrada no .env")
	}
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Banco não conectou corretamente: %v", err)
	}

	convQueries := database.New(dbConn)
	apiCfg := apiConfig{ // estpa apontando para nossa struct apiConfig
		DB: convQueries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/err", handlerErr)
	v1router.Post("/users", apiCfg.handlerCreateUser)
	v1router.Get("/users", apiCfg.handlerGetUserByAPIKey)

	router.Mount("/v1", v1router)

	srv := &http.Server{ // vai ser um pointer para o http.Server
		Handler: router,
		Addr:    ":" + portStr,
	}

	log.Printf("Server está inicializando na porta %v", portStr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
