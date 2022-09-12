package controllers

import (
	"encoding/base64"
	helpers "go-psql-gin/helpers"
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type Verifyuser struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func VerifyUser(c *gin.Context) {
	email := c.Param("email")
	password := c.Param("password")

	var encrypted_password string

	value, err := helpers.GetCachedValue(email)
	if err != nil || value == nil {

		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
		query := psql.Select("password").From("users").Where(`email = ?`, email).Where("active = ? ", true)
		sql, args, err := query.ToSql()
		if err != nil {
			log.Println(err)
			return
		}

		rows, err := dbConnect.Query(sql, args...)
		if err != nil {
			log.Printf(" error occured during insertion %v", err)
			return
		}

		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&encrypted_password)

			if err != nil {
				log.Printf("error occured during scanning the row %v", err)
				return
			}
		}
	} else {
		encrypted_password = string(value)
	}

	decrypt_password, _ := base64.StdEncoding.DecodeString(encrypted_password)
	if password == string(decrypt_password) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Successfully logged in",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  401,
			"message": "Invalid Credentials check Password",
		})
	}

}
