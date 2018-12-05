package entity

import "time"

type Weather struct {
	Lat                 string		`json:"lat" bson:"lat`
	Lon                 string    	`json:"lon" bson:"lon`
	Temp				float64	  	`json:"temp",bson:"temp""`
	Location			string		`json:"location",bson:"location"`
	Humidity			string		`json:"humidity",bson:"humidity"`
	RetrievedAt         time.Time  	`json:"created_at" bson:"created_at"`
}