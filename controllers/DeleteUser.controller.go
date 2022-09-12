package controllers

import (
	"encoding/base64"
	"go-psql-gin/domain"
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

func DeleteUser(userDetails domain.Deleteuser, c *gin.Context) {
	id := c.Param("id")
	user := userDetails
	email := user.Email
	password := base64.StdEncoding.EncodeToString([]byte(user.Password))
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Update("users").Set(`active`, false).Where("id = ? and email = ? and password = ?", id, email, password)
	sql, args, err := query.ToSql()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = dbConnect.Exec(sql, args...)

	if err != nil {
		log.Printf("Error while deleting a single user, Reason: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
}
