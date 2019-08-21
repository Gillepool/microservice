package dao

import (
	"github.com/gillepool/MovieBackend/src/movie-microservice/databases"
	"github.com/gillepool/MovieBackend/src/movie-microservice/models"
	"gopkg.in/mgo.v2/bson"
)

//MovieGenre that manages CRUD
type MovieGenre struct {
}

//GetAll gets a collection of movies
func (m *MovieGenre) GetAll() ([]models.MovieGenre, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C("genres")

	var genres []models.MovieGenre
	err := collection.Find(bson.M{}).All(&genres)
	return genres, err
}

//GetByID gets movieGenre by id
func (m *MovieGenre) GetByID(id string) (models.MovieGenre, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C("genres")

	var movieGenre models.MovieGenre
	err := collection.FindId(bson.ObjectIdHex(id)).One(&movieGenre)
	return movieGenre, err
}

//Insert movieGenre
func (m *MovieGenre) Insert(movieGenre models.MovieGenre) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C("genres")
	err := collection.Insert(&movieGenre)
	return err
}
