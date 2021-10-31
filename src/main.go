package main
import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/gonzalopages64/GO-API-REST"
)
type user struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address string `json:"address,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Zip string `json:"zip,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
type users []user

var user = users {
	{
		ID: 1,
		Firstname: "Juan",
		Lastname: "Gonzalez",
		Address: "123 Main St",
		City: "San Jose",
		State: "CA",
		Zip: "95131",
		Phone: "408-555-1234",
		Email: "test@gmail.com"
		Password: "password"

	}
}


func main() {
	router := mux.NewRouter().strictSlash(true)
	

	router.HandleFunc("/users", getUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Starting the application...")

}