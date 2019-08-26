package controller

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gillepool/MovieBackend/src/user-microservice/dao"
	"github.com/gillepool/MovieBackend/src/user-microservice/models"
	"github.com/gillepool/MovieBackend/src/user-microservice/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//User struct
type User struct {
	userDao dao.User
	utils   utils.Utils
}

func (u *User) Authenticate(ctx *gin.Context) {
	username := ctx.PostForm("user")
	password := ctx.PostForm("password")

	if _, err := u.userDao.Authenticate(username, password); err == nil {
		var tokenString string
		if tokenString, err = u.utils.GenerateJWT(username, password); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Error{Code: utils.StatusCodeUnknown, Message: err.Error()})
			log.Debug("[ERROR]: ", err)
			return
		}
		token := models.Token{Token: tokenString}
		ctx.JSON(http.StatusOK, token)

	} else {
		ctx.JSON(http.StatusUnauthorized, models.Error{Code: utils.StatusCodeUnknown, Message: err.Error()})
	}
}

func (u *User) AddUser(ctx *gin.Context) {
	var addUser models.AddUser
	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: utils.StatusCodeUnknown, Message: err.Error()})
		return
	}

	if err := addUser.Validate(); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: utils.StatusCodeUnknown, Message: err.Error()})
	}

	user := models.User{
		ID:       bson.NewObjectId(),
		Name:     addUser.Name,
		Password: addUser.Password,
	}

	if err := u.userDao.Insert(user); err == nil {
		ctx.JSON(http.StatusOK, models.Message{Message: "Successfully"})
		log.Debug("Registered a new user = " + user.Name + ", password = " + user.Password)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: utils.StatusCodeUnknown, Message: err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

func (u *User) ListUsers(ctx *gin.Context) {
	var users []models.User
	var err error
	users, err = u.userDao.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, users)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: utils.StatusCodeUnknown, Message: err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

func (u *User) GetUserByID(ctx *gin.Context) {
	var user models.User
	var err error
	id := ctx.Params.ByName("id")
	user, err = u.userDao.GetUserByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: utils.StatusCodeUnknown, Message: err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}
