package routes

import (
	snippetController "twilux/controllers/snippets"
	userController "twilux/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController    userController.UserController
	SnippetController snippetController.SnippetController
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/user")
	users.POST("/login", controller.UserController.Login)
	users.POST("/register", controller.UserController.Register)

	snippet := c.Group("/snippets")
	snippet.GET("/", controller.SnippetController.GetAll)
	snippet.POST("/", controller.SnippetController.Create)

}
