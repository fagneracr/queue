package queue

import (
	"errors"

	"go-queue/internal/system"

	"github.com/labstack/echo/v4"
)

type ReturnConfig struct {
	QueConfig interface{} `json:"configset" xml:"configset"`
	Err       []string    `json:"erros" xml:"erros" `
}

/*CreateQ - create a Queue config*/
func CreateQ() echo.HandlerFunc {
	return func(c echo.Context) error {

		var config system.Identify
		returnmsg := ReturnConfig{}
		err := c.Bind(&config)
		if err != nil {
			return err
		}
		err = system.Enq.Save(config)
		if err != nil {
			returnmsg.Err = []string{err.Error()}
		}
		returnmsg.QueConfig = config
		c.JSON(200, returnmsg)
		return errors.New("Teste")

	}

}
