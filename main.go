package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	tm := NewTodoManager()
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		todos := tm.GetAll()

		return c.JSON(http.StatusOK, todos)
	})

	authenticatedGroup := e.Group("/todos", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("authorization")
			if authorization != "auth-token" {
				c.Error(echo.ErrUnauthorized)
				return nil
			}
			next(c)
			return nil
		}
	})

	authenticatedGroup.Any("/*", func(c echo.Context) error {
		c.Error(echo.ErrNotFound)
		return nil
	})
	e.Any("/*", func(c echo.Context) error {
		c.Error(echo.ErrNotFound)
		return nil
	})
	e.Start(":8888")

}
