package main

import (
	"fmt"
	"net/http"

	"loginPage/controllers"
	"loginPage/db"

	"github.com/gorilla/mux"
)

var port = ":8000"

func main() {

	db.Users["admin@gmail.com"] = db.Userdetails{
		Name: "Admin",
		Pass: "12345",
	}

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Login).Methods("GET")
	router.HandleFunc("/", controllers.Submit).Methods("POST")
	router.HandleFunc("/home", controllers.HomePage)
	router.HandleFunc("/logout", controllers.Logout)

	router.NotFoundHandler = http.HandlerFunc(controllers.ErrorPage)

	fmt.Println("server running at ", port)
	http.ListenAndServe(port, router)

}
