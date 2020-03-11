package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lutfialfiansyah/pretest/database"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

/*GetMessageWebSocket all message when connection open.*/
func GetMessageWebSocket(c *gin.Context) {
	handler := websocket.Handler(func(conn *websocket.Conn) {
		for _, row := range database.LocalDB {
			conn.Write([]byte(row))
		}
		io.Copy(conn, conn)
	})
	handler.ServeHTTP(c.Writer, c.Request)
}

/*GetMessage all message as a array json return.*/
func GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, database.LocalDB)
	return
}