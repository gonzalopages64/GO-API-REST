package main
import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
)
type user struct {
	ID int `json:"id,omitempty"`
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
type allUsers []user

var users = allUsers {
	{
		ID: 1,
		Firstname: "Juan",
		Lastname: "Gonzalez",
		Address: "123 Main St",
		City: "San Jose",
		State: "CA",
		Zip: "95131",
		Phone: "408-555-1234",
		Email: "test@gmail.com",
		Password: "password",
	},
}

//Routes
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
	fmt.Println("Endpoint Hit: indexRoute")
}
func contactRouteGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Contact page!")
	fmt.Println("Endpoint Hit: contactRouteGET")
}
func usersRoutePOST(w http.ResponseWriter, r *http.Request) {
	
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un usurio valido")
	}
	json.Unmarshal(reqBody, &newUser)
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusCreated)	
	json.NewEncoder(w).Encode(newUser)

	fmt.Println("Endpoint Hit: usersRoutePOST")
}
func usersRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	// Parsea en json el objeto
	json.NewEncoder(w).Encode(users)

	fmt.Println("Endpoint Hit: usersRoute")
}





func main() {
	router := mux.NewRouter().StrictSlash(true)
	

	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/users", usersRoute).Methods("GET")
	router.HandleFunc("/contact", contactRouteGET).Methods("GET")
	router.HandleFunc("/users", usersRoutePOST).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Starting the application...")

}