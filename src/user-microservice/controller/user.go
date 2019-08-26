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

// Authenticate godoc
// @Summary Check user authentication
// @Description Authenticate user
// @Tags admin
// @Security ApiKeyAuth
// @Accept  multipart/form-data
// @Param user formData string true "Username"
// @Param password formData string true "Password"
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Success 200 {object} models.Token
// @Router /admin/auth [post]
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

// AddUser godoc
// @Summary Add a new user
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body models.AddUser true "Add user"
// @Failure 500 {object} models.Error
// @Failure 400 {object} models.Error
// @Success 200 {object} models.Message
// @Router /users [post]
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

// ListUsers godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 500 {object} models.Error
// @Success 200 {array} models.User
// @Router /users/list [get]
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

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.User
// @Router /users/detail/{id} [get]
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
