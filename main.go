package main

import (
	"fmt"
	config "go-psql-gin/config"
	routes "go-psql-gin/routes"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func main() {

	config.Connect()

	fmt.Println(os.Getenv("ENV"))
	if os.Getenv("ENV") == "dev" {
		envRawbytes, err := ioutil.ReadFile("config/localvariable.yaml")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		var envConfig map[string]interface{}
		err = yaml.Unmarshal(envRawbytes, &envConfig)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		for key, value := range envConfig {
			os.Setenv(key, fmt.Sprint(value))
			log.Println(key, value)
		}

	}

	router := gin.Default()

	routes.Routes(router)

	config.RedisConnect()

	log.Fatal(router.Run(":8080"))
}
