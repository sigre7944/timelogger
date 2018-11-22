package main

import (
	"encoding/json"
	. "github.com/user/timelogger2/timelog"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"log"
	//"math/rand"
	"net/http"

	//"strconv"

	"github.com/gorilla/mux"
)

type Repository struct{}
type Controller struct {
	Repository Repository
}

var controller = &Controller{Repository: Repository{}}

const SERVER = "localhost:27017"
const DBNAME = "timelogStore"
const DOCNAME = "timelog"

//Init variable
var timelogs []Timelog

func AllTimelogEndPoint(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	var results []Timelog
	err = c.Find(nil).All(&results)
	if err != nil {
		// TODO: Do something about the error
	} else {
		fmt.Println("Results All: ", results)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func CreateTimelogEndPoint(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var timelog Timelog
	_ = json.NewDecoder(r.Body).Decode(&timelog)
	timelog.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(timelog)

	if err != nil {
		log.Fatal(err)
	}

	//respondWithJson(w, http.StatusCreated, timelog)
	json.NewEncoder(w).Encode(timelog)
}

func FindTimelogByDate(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	var results []Timelog

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get parameters
	err = c.Find(bson.M{"day": params["day"]}).All(&results)
	// for _, item := range timelogs {
	// 	if item.Day == params["day"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	json.NewEncoder(w).Encode(results)
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
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

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
