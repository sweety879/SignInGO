package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-psql-gin/domain"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExternalApi(user domain.ExternalAPILogin, c *gin.Context) {

	newUser := user
	fmt.Println(newUser)
	jsonReq, err := json.Marshal(newUser)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(jsonReq)
	fmt.Println(string(jsonReq))
	resp, err := http.Post("https://report-api.dotnu.co/login", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Println(bodyString)

	var loginStruct domain.ExternalAPILogin
	json.Unmarshal(bodyBytes, &loginStruct)

	if loginStruct.UserName == newUser.UserName {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "successfully logged in",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong credentials"})
	}

}
