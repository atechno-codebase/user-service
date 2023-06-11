package server

import (
	"ates/services/user-service/config"
	"ates/services/user-service/database"
	"ates/services/user-service/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	claims := ctx.Value("claims").(map[string]any)
	_username, ok := claims["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewErrorResponse("access_token does not contain `username`").ByteResponse())
		return
	}

	username, ok := _username.(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewErrorResponse("`username` from access_token is not a string").ByteResponse())
		return
	}
	log.Println(username)

	var users []models.User

	_, err := database.RunQuery(func(client *mongo.Client) (interface{}, error) {
		collection := client.Database(config.Configuration.Database).Collection("user")

		cursor, err := collection.Find(ctx, bson.M{"createdBy": username})
		if err != nil {
			return nil, err
		}
		err = cursor.All(ctx, &users)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewErrorResponse("error while readings users from database").ByteResponse())
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewErrorResponse("error while marshalling users").ByteResponse())
		return
	}

	w.WriteHeader(200)
	w.Write(response)
	return
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dataBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewErrorResponse("error while reading request body").ByteResponse())
		return
	}

	postData := gjson.ParseBytes(dataBytes)

	var user models.User
	_, err = database.RunQuery(func(client *mongo.Client) (interface{}, error) {
		collection := client.Database(config.Configuration.Database).Collection("user")

		res := collection.FindOne(ctx, bson.M{
			"username": postData.Get("username").String(),
		})
		if res.Err() != nil {
			return nil, err
		}
		err := res.Decode(&user)
		if err != nil {
			return nil, err
		}
		return res, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewErrorResponse("error while reading user from database").ByteResponse())
		return
	}

	if user.Verification.Code == 

	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewErrorResponse("error while marshalling users").ByteResponse())
		return
	}

	w.WriteHeader(200)
	w.Write(response)
	return

}

func LoginUser(w http.ResponseWriter, r *http.Request) {}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {}

func ResetPassword(w http.ResponseWriter, r *http.Request) {}

func GetProfile(w http.ResponseWriter, r *http.Request) {}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
