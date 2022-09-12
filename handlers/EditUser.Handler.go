package handlers

import (
	controllers "go-psql-gin/controllers"
	"go-psql-gin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditUserHandler(c *gin.Context) {

	var oldData domain.EditUserType
	if err := c.ShouldBindJSON(&oldData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controllers.EditUser(oldData, c)

}
