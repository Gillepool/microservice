package dao

import (
	"github.com/gillepool/MovieBackend/src/movie-microservice/databases"
	"github.com/gillepool/MovieBackend/src/movie-microservice/models"
	"gopkg.in/mgo.v2/bson"
)

//Movie that manages CRUD
type Movie struct {
}

//GetAll gets a collection of movies
func (m *Movie) GetAll() ([]models.Movie, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C("movies")

	var movies []models.Movie
	err := collection.Find(bson.M{}).All(&movies)
	return movies, err
}

//GetByID gets movie by id
func (m *Movie) GetByID(id string) (models.Movie, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C("movies")

	var movie models.Movie
	err := collection.FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

//Insert movie
func (m *Movie) Insert(movie models.Movie) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C("movies")
	err := collection.Insert(&movie)
	return err
}
