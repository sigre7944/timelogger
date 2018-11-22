package dao

import (
	"log"

	. "github.com/user/timelogger2/timelog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TimelogDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "timelog"
)

func (t *TimelogDAO) Connect() {
	session, err := mgo.Dial(t.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(t.Database)
}

func (t *TimelogDAO) FindAll() ([]Timelog, error) {
	var timelogs []Timelog
	err := db.C(COLLECTION).Find(bson.M{}).All(&timelogs)
	return timelogs, err
}

func (t *TimelogDAO) FindByDate(date string) (Timelog, error) {
	var timelog Timelog
	err := db.C(COLLECTION).Find(bson.M{"day": date}).All(&timelog)
	return timelog, err
}

func (t *TimelogDAO) Insert(timelog Timelog) error {
	err := db.C(COLLECTION).Insert(&timelog)
	return err
}
