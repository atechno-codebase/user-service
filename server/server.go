package server

import (
	"ates/services/user-service/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	PORT := config.Get("port").String()
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

	/*
		r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			methods, _ := route.GetMethods()
			pre, _ := route.GetPathRegexp()
			pre = strings.ReplaceAll(pre, "^", "")
			pre = strings.ReplaceAll(pre, "$", "")
			fmt.Printf("\n%+v %v\n", methods[0], pre)
			return nil
		})
	*/

}
