package main

import (
	"fmt"
	"go-queue/internal/grpcserve"
	"go-queue/internal/qsys"
	"go-queue/internal/queue"
	"net/http"
	"path"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {

	err := qsys.InitQ(qsys.Conf{
		Directory: path.Join("/home/fribeiro/fagner/qsys"),
	})
	if err != nil {
		fmt.Println("Error trying initialize system: " + err.Error())
	}
	for x := 0; x < 2; x++ {
		name := "fila" + strconv.Itoa(x)
		if !qsys.ExistsQ(name) {
			err = qsys.CreateQ(qsys.ConfigQueue{
				Name:       name,
				Persistent: true,
				Variable: []qsys.Variable{{Key: "teste",
					Value: "teste1"}},
			})
			if err != nil {
				fmt.Println("Error trying initialize system: " + err.Error())
			}
			fmt.Println("Queue created: " + name)
		}
	}
	listq := qsys.ListQ()
	fmt.Println(listq)
	err = qsys.DeleteQ("fila0")
	if err != nil {
		fmt.Println(err)
	}
	//Initiate Grpc Server
	go grpcserve.InitServer()
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
