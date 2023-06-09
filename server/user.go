package server

import (
	"ates/services/user-service/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while reading request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrorToResponse(err))
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("error while reading request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrorToResponse(err))
		return
	}

	now := time.Now().Unix()
	fmt.Printf("%+v\n", user)

	claims := r.Context().Value("claims").(map[string]any)

	rand.Seed(now)
	fmt.Println(claims)
	user.CreatedBy = claims["username"].(string)
	user.CreatedOn = now
	user.Verification.IsVerified = false
	user.Verification = models.Verification{
		IsVerified: false,
		Code:       strconv.Itoa(rand.Intn(9999-1000) + 1000),
	}

	fmt.Println()

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {}

func VerifyUser(w http.ResponseWriter, r *http.Request) {}

func LoginUser(w http.ResponseWriter, r *http.Request) {}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {}

func ResetPassword(w http.ResponseWriter, r *http.Request) {}

func GetProfile(w http.ResponseWriter, r *http.Request) {}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
