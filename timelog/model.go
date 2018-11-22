package timelog

import "gopkg.in/mgo.v2/bson"

type Timelog struct {
	ID bson.ObjectId `json:"_id" binding:"required"`
	// Year      int    `json:"year"`
	// Month     int    `json:"month"`
	Day string `json:"day" binding:"required"`
	// Hour      int    `json:"hour"`
	// Minute    int    `json:"minute"`
	// Second    int    `json:"second"`
	PointType string `json:"pointtype" binding:"required"`
	Activity  string `json:"activity" binding:"required"`
}
