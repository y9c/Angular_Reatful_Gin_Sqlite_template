/*********************************************************************************
*     File Name           :     api.go
*     Created By          :     yc
*     Creation Date       :     [2018-02-24 17:00]
*     Last Modified       :     [2018-02-24 20:37]
*     Description         :
**********************************************************************************/

package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// people api include these function:
// InitPeople GetPeople GetPerson CreatePerson DeletePerson

// main function to boot up everything
func API() {
	InitPeople()
	router := mux.NewRouter()
	router.HandleFunc("/api/people", GetPeople).Methods("GET")
	router.HandleFunc("/api/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/api/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/api/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
