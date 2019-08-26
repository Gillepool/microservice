package main

import (
	"github.com/gillepool/MovieBackend/src/user-microservice/controller"
	"github.com/gillepool/MovieBackend/src/user-microservice/databases"
	"github.com/gillepool/MovieBackend/src/user-microservice/utils"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	if err = utils.LoadConfig(); err != nil {
		return err
	}

	m.router = gin.Default()

	return err
}

func main() {

	m := Main{}

	if m.initServer != nil {
		return
	}

	defer databases.Database.Close()

	c := controller.User{}

	v1 := m.router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Authenticate)
		}

		user := v1.Group("/users")

		// APIs need to be authenticated
		user.Use(jwt.Auth(utils.Config.JwtSecretPassword))
		{
			user.POST("", c.AddUser)
			user.GET("/list", c.ListUsers)
			user.GET("detail/:id", c.GetUserByID)
		}
	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(utils.Config.Port)

}
