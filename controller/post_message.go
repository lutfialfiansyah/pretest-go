package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lutfialfiansyah/pretest/database"
	"net/http"
)

// PostMassage used for save massage. All message will save into array.
func PostMassage(c *gin.Context) {
	message := c.Param("message")

	database.LocalDB = append(database.LocalDB, message)

	c.String(http.StatusOK, message+" Successfully")

	return
}