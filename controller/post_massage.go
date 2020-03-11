package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lutfialfiansyah/pretest/database"
	"net/http"
)

/*SendMessage use for save message. All message will save into array.*/
func SendMessage(c *gin.Context) {
	message := c.Param("message")
	database.LocalDB = append(database.LocalDB, message)
	c.String(http.StatusOK, message+" successfully")
	return
}
