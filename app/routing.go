package app

import (
	"github.com/lutfialfiansyah/pretest/controller"
)

// mapUrl listed all mapping URLs.

func mapUrl() {

	router.GET("/all-message", controller.GetMessageWS)
	router.POST("/post/:message", controller.PostMassage)
	router.GET("/get/all-message", controller.GetMessage)

}
