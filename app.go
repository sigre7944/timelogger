package main

import (
	"encoding/json"
	"fmt"
	. "github.com/user/timelogger2/dao"
	. "github.com/user/timelogger2/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dao = TimelogDAO{}

func AllTimelogEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateTimelogEndPoint(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close()
	// var timelog Timelog
	// //decode the request body into a timelog object, assign it an ID, and use DAO insert to create a timelog in the database
	// if err := json.NewDecoder(r.Body).Decode(&timelog); err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	// 	return
	// }
	// timelog.ID = bson.NewObjectId()
	// if err := dao.Insert(timelog); err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// respondWithJson(w, http.StatusCreated, timelog)
}

func FindTimelogByDate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/timelog", AllTimelogEndPoint).Methods("GET")
	r.HandleFunc("/timelog", CreateTimelogEndPoint).Methods("POST")
	r.HandleFunc("/timelog/{day}", FindTimelogByDate).Methods("GET")
	if err := http.ListenAndServe(":3002", r); err != nil {
		log.Fatal(err)
	}
}
