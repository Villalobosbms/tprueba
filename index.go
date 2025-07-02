package main

import (
	"net/http"

	"git.hub/villalobosbms/tprueba/routes"
	"github.com/gorilla/mux"
)

func main(){
	route := mux.NewRouter()

	route.HandleFunc("/", routes.Index)

	route.HandleFunc("/user", routes.GetUsers).Methods("GET")
	route.HandleFunc("/user/{id}", routes.GetUser).Methods("GET")
	route.HandleFunc("/user/{id}", routes.CreateUser).Methods("POST")
	route.HandleFunc("/user/{id}", routes.UpdateUser).Methods("PUT")
	route.HandleFunc("/user/{id}", routes.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":3000", route)
}