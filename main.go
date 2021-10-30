package main
import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func getPeople ( w http.ResponseWriter, req *http.Request ) {
}

func main() {
	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/", Index)
	router.HandleFunc("/tasks", TasksIndex)
	router.HandleFunc("/tasks/{id}", TasksShow)
	log.Fatal(http.ListenAndServe(":8000", router))
}