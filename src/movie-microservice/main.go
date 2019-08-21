package main

import (
	"github.com/gillepool/MovieBackend/src/movie-microservice/controllers"
	"github.com/gillepool/MovieBackend/src/movie-microservice/databases"
	"github.com/gillepool/MovieBackend/src/movie-microservice/utils"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) initServer() error {
	var err error
	err = utils.LoadConfig()
	if err != nil {
		return err
	}

	err = databases.Database.Init()
	if err != nil {
		return err
	}

	r.router = gin.Default()
	return nil
}

func main() {
	r := Router{}
	if r.initServer() != nil {
		return
	}

	defer databases.Database.Close()

	cMovie := controllers.Movie{}
	cGenre := controllers.MovieGenre{}

	// Simple group: v1
	v1 := r.router.Group("/api/v1")
	{
		v1.GET("/movies/list", cMovie.ListMovies)
		v1.GET("/movies/list/:id", cMovie.ListMovieByID)

		v1.GET("/genres/list", cGenre.ListMovieGenres)
		v1.GET("/genres/list/:id", cGenre.ListMovieGenreByID)

		// APIs need to use token string
		v1.POST("/movies", cMovie.AddMovie)
		v1.POST("/genres", cGenre.AddMovieGenre)

	}

	//m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.router.Run(":8080")
}
