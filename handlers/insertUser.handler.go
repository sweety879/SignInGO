package handlers

import (
	"net/http"
	"regexp"

	controllers "go-psql-gin/controllers"
	domain "go-psql-gin/domain"

	"github.com/gin-gonic/gin"
)

var regex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func InsertUserHandler(c *gin.Context) {

	var newUser domain.Insertuser

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !regex.MatchString(newUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email field should be a valid email address!"})
		return
	}

	controllers.InsertUser(newUser, c)

}
