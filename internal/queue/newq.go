package queue

type ReturnConfig struct {
	QueConfig interface{} `json:"configset" xml:"configset"`
	Err       []string    `json:"erros" xml:"erros" `
}

// /*CreateQ - create a Queue config*/
// func (s *system) CreateQ() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		var config system.Identify
// 		returnmsg := ReturnConfig{}
// 		err := c.Bind(&config)
// 		if err != nil {
// 			return err
// 		}
// 		err = system.Save(config)
// 		if err != nil {
// 			returnmsg.Err = []string{err.Error()}
// 		}
// 		returnmsg.QueConfig = config
// 		c.JSON(200, returnmsg)
// 		return errors.New("Teste")

// 	}

// }
