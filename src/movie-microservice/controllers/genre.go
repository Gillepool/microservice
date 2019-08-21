package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gillepool/MovieBackend/src/movie-microservice/dao"
	"github.com/gillepool/MovieBackend/src/movie-microservice/models"
	"github.com/gin-gonic/gin"
)

//MovieGenre information
type MovieGenre struct {
	movieGenreDAO dao.MovieGenre
}

//AddMovieGenre to database
func (m *MovieGenre) AddMovieGenre(ctx *gin.Context) {
	var movieGenre models.MovieGenre
	if err := ctx.BindJSON(&movieGenre); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	movieGenre.ID = bson.NewObjectId()
	err := m.movieGenreDAO.Insert(movieGenre)

	if err == nil {
		fmt.Println(http.StatusOK)
	} else {
		fmt.Println("Error")
	}
}

//ListMovieGenres lists all movies
func (m *MovieGenre) ListMovieGenres(ctx *gin.Context) {
	var movieGenres []models.MovieGenre
	var err error
	movieGenres, err = m.movieGenreDAO.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, movieGenres)
	} else {
		fmt.Println("Cannot retrive movies right now")
	}
}

//ListMovieGenreByID lists all movies
func (m *MovieGenre) ListMovieGenreByID(ctx *gin.Context) {
	var movieGenre models.MovieGenre
	var err error
	id := ctx.Params.ByName("id")
	movieGenre, err = m.movieGenreDAO.GetByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, movieGenre)
	} else {
		fmt.Println("Cannot retrive movie right now")
	}
}
