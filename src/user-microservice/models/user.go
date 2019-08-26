package models

import (
	"gopkg.in/mgo.v2/bson"
)

//User struct
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name     string        `bson:"name" json:"name" example:"gman"`
	Password string        `bson:"password" json:"password" example:"gman"`
}

type AddUser struct {
	Name     string `bson:"name" json:"name" example:"gman"`
	Password string `bson:"password" json:"password" example:"gman"`
}

//Validate user
func (u AddUser) Validate() error {
	var err error
	switch {
	case len(u.Name) == 0:
		return err
	default:
		return nil
	}
}
