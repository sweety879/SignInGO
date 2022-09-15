package controllers

import (
	"encoding/base64"
	"encoding/json"

	"go-psql-gin/domain"
	helpers "go-psql-gin/helpers"
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

func EditUser(user domain.EditUserType, c *gin.Context) {
	userId := c.Param("id")
	oldData := user

	password := base64.StdEncoding.EncodeToString([]byte(oldData.NewPassword))
	oldData.OldPassword = base64.StdEncoding.EncodeToString([]byte(oldData.OldPassword))

	var newData domain.Insertuser
	newData.Email = oldData.Email
	newData.Password = password

	jsonReq, jsonErr := json.Marshal(newData)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
		return
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Update("users").Set("password", password).Set("info", string(jsonReq)).Where("id = ? and email =? and password = ? and active = ?", userId, oldData.Email, oldData.OldPassword, true)
	sql, args, err := query.ToSql()

	if err != nil {
		log.Println(err)
		return
	}

	res, err := dbConnect.Exec(sql, args...)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "server error",
		})
		return
	} else {
		count, err := res.RowsAffected()
		if err != nil {
			log.Println(err.Error())
		} else {
			if count != 0 {
				err = helpers.EnCache(oldData.Email, []byte(password))
				if err != nil {
					log.Printf("redis enchanche error %v", err)
					log.Println()
				}
				c.JSON(http.StatusOK, gin.H{
					"status":  200,
					"message": "User Edited Successfully",
				})

			} else {

				c.JSON(http.StatusOK, gin.H{
					"status":  200,
					"message": "Invalid Credentials or account has been deleted",
				})
			}

		}
	}

}
