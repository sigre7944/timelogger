package models

type Timelog struct {
	ID string `json:"_id" binding:"required"`
	// Year      int    `json:"year"`
	// Month     int    `json:"month"`
	Day string `json:"day" binding:"required"`
	// Hour      int    `json:"hour"`
	// Minute    int    `json:"minute"`
	// Second    int    `json:"second"`
	PointType string `json:"pointtype" binding:"required"`
	Activity  string `json:"activity" binding:"required"`
}
