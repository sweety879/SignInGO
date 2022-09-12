package handlers

import (
	"go-psql-gin/controllers"
	"go-psql-gin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExternalApiHandler(c *gin.Context) {

	var user domain.ExternalAPILogin

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controllers.ExternalApi(user, c)

}
