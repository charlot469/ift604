package merge

import (
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

var message []map[string]interface{}
func GitlabHook(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for{
			var results []map[string]interface{}
			err := websocket.Message.Receive(ws, results)

			if err != nil {
				return
			}

			message = results
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func Notification(c echo.Context) error{
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for{
			if message != nil{
				sendError := websocket.Message.Send(ws, message)

				if sendError != nil {
					c.Logger().Error(sendError)
					return
				}

				message = nil
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

