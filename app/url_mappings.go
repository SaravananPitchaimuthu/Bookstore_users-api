package app

import (
	"github.com/SaravananPitchaimuthu/Bookstore_users-api/controllers/users"
)

func mapURLs() {
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
