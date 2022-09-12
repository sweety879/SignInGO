package config

import (
	"log"
	"os"

	helpers "go-psql-gin/helpers"

	"github.com/go-redis/redis"
)

// Connecting to db
func RedisConnect() *redis.Client {

	redisHostPortString := os.Getenv("redisHostPortString")
	redisUrl, err := redis.ParseURL(redisHostPortString)
	if err != nil {
		log.Println(err, "Initializing Redis Failed Due to fetching url")
		os.Exit(1)
	}

	redisOptions := redis.Options{
		Addr: redisUrl.Addr,
		DB:   redisUrl.DB,
	}
	log.Println("redisOptions", "Initializing Redis- RedisOptions")

	client := redis.NewClient(&redis.Options{
		Addr:     redisOptions.Addr,
		Password: "",
		DB:       redisOptions.DB,
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)

	helpers.InitiateRDS(client)

	return client

}
