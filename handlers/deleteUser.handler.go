package handlers

import (
	"go-psql-gin/controllers"
	"go-psql-gin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(c *gin.Context) {

	var user domain.Deleteuser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !regex.MatchString(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email field should be a valid email address!"})
		return
	}

	controllers.DeleteUser(user, c)

}
