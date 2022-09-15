package controllers

import (
	"encoding/base64"
	"encoding/json"

	helpers "go-psql-gin/helpers"
	"log"
	"net/http"

	domain "go-psql-gin/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

func InsertUser(user domain.Insertuser, c *gin.Context) {

	var Id string
	newUser := user
	email := newUser.Email
	password := base64.StdEncoding.EncodeToString([]byte(newUser.Password))
	newUser.Password = password
	jsonReq, jsonErr := json.Marshal(newUser)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
		return
	}
	// fmt.Println(string(jsonReq))
	query := sq.Insert("users").Columns("email", "password", "info").Values(email, password, string(jsonReq)).Suffix("RETURNING \"id\"").RunWith(dbConnect).PlaceholderFormat(sq.Dollar)
	query.QueryRow().Scan(&Id)

	// fmt.Println(email)
	// fmt.Println(password)
	err := helpers.EnCache(email, []byte(password))
	if err != nil {
		log.Printf("redis enchanche error %v", err)
		log.Println()
	}

	log.Printf("successfullyy insert a row with id %v", Id)

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "successfully inserted the data",
	})
}
