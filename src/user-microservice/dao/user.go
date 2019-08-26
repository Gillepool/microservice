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

//GetUserByID returns the user of id
func (u *User) GetUserByID(id string) (models.User, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(utils.ColUsers)

	var user models.User
	err := collection.FindId(bson.IsObjectIdHex(id)).One(&user)

	return user, err
}

//DeleteUser removes a user from the collection
func (u *User) DeleteUser(user User) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(utils.ColUsers)
	err := collection.Remove(&user)

	return err
}

//Authenticate authenticates user credentials
func (u *User) Authenticate(name string, password string) (models.User, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(utils.ColUsers)
	var user models.User
	err := collection.Find(bson.M{"$and": []bson.M{bson.M{"name": name}, bson.M{"password": password}}}).One(&user)
	return user, err
}

func (u *User) Insert(user models.User) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(utils.ColUsers)

	err := collection.Insert(&user)
	return err
}
