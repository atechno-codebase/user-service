package server

import (
	"ates/services/user-service/config"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

const (
	AUTH_HEADER = "Authorization"
)

var (
	ErrExtractClaims = errors.New("could not extract claims")
	ErrJWTParse      = errors.New("could not parse jwt")
	ErrNoAuthHeader  = errors.New("invalid value for header: 'Authorization'")
	ErrNoAuthToken   = errors.New("token missing in 'Authorization' header")
)

func VerifyToken(token, secretToken string) (map[string]interface{}, error) {
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrJWTParse
		}
		return []byte(secretToken), nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if ok && tok.Valid {
		return claims, nil
	}

	return nil, ErrExtractClaims
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(AUTH_HEADER)
		if authHeader == "" {
			err := ErrNoAuthHeader
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ErrorToResponse(err))
			return
		}

		bearer_n_token := strings.Split(authHeader, " ")
		if len(bearer_n_token) < 2 {
			err := ErrNoAuthToken
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ErrorToResponse(err))
			return
		}
		jwtToken := bearer_n_token[1]
		if jwtToken == "" {
			err := ErrNoAuthToken
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ErrorToResponse(err))
			return
		}

		claims, err := VerifyToken(jwtToken, config.Get("secretToken").String())
		if err != nil {
			log.Println("error while decoding jwt token:", err)
			http.Error(w, err.Error(), http.StatusInsufficientStorage)
		}

		newReq := r.WithContext(
			context.WithValue(r.Context(), "claims", claims),
		)
		next(w, newReq)
	}
}
