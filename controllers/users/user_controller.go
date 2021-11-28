package users

import (
	"net/http"

	"github.com/SaravananPitchaimuthu/Bookstore_users-api/domain/users"
	"github.com/SaravananPitchaimuthu/Bookstore_users-api/services"
	"github.com/SaravananPitchaimuthu/Bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me !!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		// handle bad request
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		// handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me !!")
}
