package queue

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"go-queue/internal/system"
	"go-queue/internal/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//	type QueConfig struct {
//		Name  string `json:"name" xml:"name"`
//		Type  string `json:"type" xml:"type"`
//		Limit int    `json:"limit,omitempty" xml:"limit,omitempty"`
//		TTL   int    `json:"ttl,omitempty" xml:"tt,omitempty"`
//	}
type ReturnConfig struct {
	QueConfig interface{} `json:"configset" xml:"configset"`
	Err       []string    `json:"erros" xml:"erros" `
}

/*CreateQ - create a Queue config*/
func CreateQ() echo.HandlerFunc {
	return func(c echo.Context) error {

		var config map[string]interface{}
		returnmsg := ReturnConfig{}
		err := c.Bind(&config)
		if err != nil {
			return err
		}
		if filepath.IsLocal("./config") {
			filename := filepath.Join("./config", config["name"].(string)+".conf")
			if utils.FileExists(filename) {
				returnmsg.QueConfig = config
				returnmsg.Err = []string{"Configuration name alredy created"}
				c.JSON(500, returnmsg)
				return nil
			}
			var er error
			f, err := os.Create(filename)
			if err != nil {
				er = errors.Join(err, er)

			}
			defer f.Close()
			config["id"] = uuid.New()
			identString, err := json.MarshalIndent(config, " ", " ")
			if err != nil {
				er = errors.Join(err, er)

			}
			_, err = f.WriteString(string(identString))
			system.Enq.Createq(config["id"].(uuid.UUID))
			er = errors.Join(err, er)
			if er != nil {
				eacherr := strings.Split(er.Error(), "\n")
				returnmsg.QueConfig = config
				returnmsg.Err = eacherr
				c.JSON(500, returnmsg)
				return er
			}

		}
		c.JSON(200, config)
		return errors.New("Teste")

	}

}
