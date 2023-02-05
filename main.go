package main

import (
	"fmt"
	"go-queue/internal/queue"
	"go-queue/internal/system"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	q, err := system.InitQ()
	if err != nil {
		fmt.Println("Error trying initialize system: " + err.Error())
	}

	err = q.Createq(system.ConfigQueue{
		Name:       "fila1",
		Persistent: true,
		Variable: []system.Variable{{Key: "teste",
			Value: "teste1"}},
	})
	if err != nil {
		fmt.Println("Error trying initialize system: " + err.Error())
	}

	e := echo.New()
	//e.POST("/qcreate", queue.CreateQ())
	e.POST("/queue/:name", queue.Enquete())
	e.Logger.Fatal(e.Start(":8080"))
}
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
