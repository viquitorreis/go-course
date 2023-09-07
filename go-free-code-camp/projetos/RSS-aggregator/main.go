package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	portStr := os.Getenv("PORT")

	if portStr == "" {
		log.Fatal("PORTA não encontrada no .env") // vai parar o programa imediatamente com codigo 1 e uma mensagem
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

	router.Mount("/v1", v1router)

	srv := &http.Server{ // vai ser um pointer para o http.Server
		Handler: router,
		Addr:    ":" + portStr,
	}

	log.Printf("Server está inicializando na porta %v", portStr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
