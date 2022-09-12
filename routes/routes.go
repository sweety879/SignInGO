package routes

import (
	controllers "go-psql-gin/controllers"
	handlers "go-psql-gin/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/user", handlers.InsertUserHandler)
	router.GET("/user/:id", controllers.GetUser)
	router.PUT("/user/delete/:id", handlers.DeleteUserHandler)
	router.PATCH("/user/update/:id", handlers.EditUserHandler)
	router.GET("/user/login/:email/:password", controllers.VerifyUser)
}
