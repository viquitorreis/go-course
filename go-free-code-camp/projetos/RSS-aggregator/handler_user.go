package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gitlab.com/victorreisprog/go/-/tree/master/go-free-code-camp/projetos/RSS-aggregator/internal/auth"
	"gitlab.com/victorreisprog/go/-/tree/master/go-free-code-camp/projetos/RSS-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Erro ao passar o JSON")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, "Não conseguiu criar o user")
		return
	}

	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, "Erro no Auth -> handler do user")
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, "Não conseguiu achar a APIKey do user no DB")
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}