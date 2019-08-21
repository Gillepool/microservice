package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gillepool/MovieBackend/src/movie-microservice/dao"
	"github.com/gillepool/MovieBackend/src/movie-microservice/models"
	"github.com/gin-gonic/gin"
)

type Movie struct {
	movieDAO dao.Movie
}

//AddMovie to database
func (m *Movie) AddMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	movie.ID = bson.NewObjectId()
	err := m.movieDAO.Insert(movie)

	if err == nil {
		fmt.Println(http.StatusOK)
	} else {
		fmt.Println("Error")
	}
}

//ListMovies lists all movies
func (m *Movie) ListMovies(ctx *gin.Context) {
	var movies []models.Movie
	var err error
	movies, err = m.movieDAO.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, movies)
	} else {
		fmt.Println("Cannot retrive movies right now")
	}
}

//ListMovieByID lists all movies
func (m *Movie) ListMovieByID(ctx *gin.Context) {
	var movie models.Movie
	var err error
	id := ctx.Params.ByName("id")
	movie, err = m.movieDAO.GetByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, movie)
	} else {
		fmt.Println("Cannot retrive movie right now")
	}
}
