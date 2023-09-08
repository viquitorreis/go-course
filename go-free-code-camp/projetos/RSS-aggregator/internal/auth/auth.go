package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("nenhuma autenticação encontrada")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 { // esperamos que o valor das keys seja 2 valores separados pro espaço
		return "", errors.New("header de autenticação mal formado")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("começo do Header de autenticação mal formado")
	}

	return vals[1], nil
}
