package entity

import "time"

type Weather struct {
	Lat                 string		`json:"lat" bson:"lat`
	Lon                 string    	`json:"lon" bson:"lon`
	Temp				float64	  	`json:"temp",bson:"temp""`
	Sunrise				int64	  	`json:"sunrise",bson:"sunrise""`
	Sunset				int64	  	`json:"sunset",bson:"sunset""`
	Location			string		`json:"location",bson:"location"`
	Humidity			float64		`json:"humidity",bson:"humidity"`
	RetrievedAt         time.Time  	`json:"created_at" bson:"created_at"`
}