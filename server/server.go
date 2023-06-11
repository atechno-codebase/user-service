package server

import (
	"ates/services/user-service/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	PORT := config.Configuration.Port
	r := mux.NewRouter()
	initRouter(r)

	log.Printf("started server on port: %s\n", PORT)
	log.Fatalln(http.ListenAndServe(PORT, r))
}

func initRouter(r *mux.Router) {

	r.Handle("/register", authMiddleware(RegisterUser)).Methods(http.MethodPost)

	r.Handle("/allusers", authMiddleware(GetAllUsers)).Methods(http.MethodGet)

	r.HandleFunc("/verify", VerifyUser).Methods(http.MethodPost)

	r.HandleFunc("/login", LoginUser).Methods(http.MethodPost)

	r.HandleFunc("/forgot", ForgotPassword).Methods(http.MethodPost)

	r.HandleFunc("/resetpwd", ResetPassword).Methods(http.MethodPost)

	r.HandleFunc("/delete", DeleteUser).Methods(http.MethodDelete)

	r.Handle("/about", authMiddleware(GetProfile)).Methods(http.MethodGet)

}
