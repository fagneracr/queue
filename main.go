package main

import (
	"fmt"
	"go-queue/internal/queue"
	"go-queue/internal/system"
	"net/http"
	"path"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {

	q, err := system.InitQ(system.Conf{
		Directory: path.Join("/home/fribeiro/fagner/qsys"),
	})
	if err != nil {
		fmt.Println("Error trying initialize system: " + err.Error())
	}
	for x := 0; x < 2; x++ {
		name := "fila" + strconv.Itoa(x)
		if !q.ExistsQ(name) {
			err = q.CreateQ(system.ConfigQueue{
				Name:       name,
				Persistent: true,
				Variable: []system.Variable{{Key: "teste",
					Value: "teste1"}},
			})
			if err != nil {
				fmt.Println("Error trying initialize system: " + err.Error())
			}
			fmt.Println("Queue created: " + name)
		}
	}
	listq := q.ListQ()
	fmt.Println(listq)
	err = q.DeleteQ("fila0")
	if err != nil {
		fmt.Println(err)
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
