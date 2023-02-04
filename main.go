package main

import (
	"go-queue/internal/queue"
	"go-queue/internal/system"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	system.InitEnq()
	e := echo.New()
	e.POST("/qcreate", queue.CreateQ())
	e.POST("/queue/:name", queue.Enquete())
	e.Logger.Fatal(e.Start(":8080"))
}
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
