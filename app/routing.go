package app

import (
	"github.com/lutfialfiansyah/pretest/controller"
)

// to listen all mapping URLs.
func mapUrls() {
	router.GET("/get-message", controller.GetMessage)
	router.GET("/all-message", controller.GetMessageWebSocket)
	router.POST("/post/:message", controller.SendMessage)
}
