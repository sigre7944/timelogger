package main

import (
	"encoding/json"
	. "github.com/user/timelogger2/dao"
	. "github.com/user/timelogger2/models"

	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Init variable
var timelogs []Timelog

var dao = TimelogDAO{}

func AllTimelogEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timelogs)
}

func CreateTimelogEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var timelog Timelog
	_ = json.NewDecoder(r.Body).Decode(&timelog)
	timelog.ID = strconv.Itoa(rand.Intn(1000000))
	timelogs = append(timelogs, timelog)
	json.NewEncoder(w).Encode(&Timelog(timelog))
}

func FindTimelogByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get parameters

	for _, item := range timelogs {
		if item.Day == params["day"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Timelog{})
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

	// Mock Data -@todo -implement DB
	timelogs = append(timelogs, Timelog{ID: "1", Day: "21", PointType: "start", Activity: "Play"})
	timelogs = append(timelogs, Timelog{ID: "2", Day: "22", PointType: "stop", Activity: "Play"})

	r.HandleFunc("/timelog", AllTimelogEndPoint).Methods("GET")
	r.HandleFunc("/timelog", CreateTimelogEndPoint).Methods("POST")
	r.HandleFunc("/timelog/{day}", FindTimelogByDate).Methods("GET")

	if err := http.ListenAndServe(":3002", r); err != nil {
		log.Fatal(err)
	}
}
