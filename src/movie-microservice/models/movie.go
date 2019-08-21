package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Movie information
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	URL         string        `bson:"url" json:"url"`
	CoverImage  string        `bson:"coverImage" json:"coverImage"`
	Description string        `bson:"description" json:"description"`
	Genre       []MovieGenre  `bson:"genre" json:"genre"`
}
