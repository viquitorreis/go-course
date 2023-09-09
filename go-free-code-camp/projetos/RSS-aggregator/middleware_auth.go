package main //vai ser parte do pacote main

import (
	"fmt"
	"net/http"

	"gitlab.com/victorreisprog/go/-/tree/master/go-free-code-camp/projetos/RSS-aggregator/internal/auth"
	"gitlab.com/victorreisprog/go/-/tree/master/go-free-code-camp/projetos/RSS-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Erro no Auth -> handler do user %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Não conseguiu achar a APIKey do user no DB %v", err))
			return
		}

		handler(w, r, user)
	}
}
