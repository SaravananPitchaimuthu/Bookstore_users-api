package app

import (
	"github.com/SaravananPitchaimuthu/Bookstore_users-api/controllers/users"
)

func mapURLs() {
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("internal/users/search", users.Search)

}
