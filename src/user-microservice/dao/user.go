package dao

import (
	"github.com/gillepool/MovieBackend/src/user-microservice/databases"
	"github.com/gillepool/MovieBackend/src/user-microservice/models"
	"github.com/gillepool/MovieBackend/src/user-microservice/utils"

	"gopkg.in/mgo.v2/bson"
)

//User struct
type User struct {
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(utils.ColUsers)

	var users []models.User
	err := collection.Find(bson.M{}).All(&users)
	return users, err
}
