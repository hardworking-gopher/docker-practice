package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	redisHostNameKey       = "REDIS_HOST_NAME"
	redisVisitsKey         = "visits"
	internalServerErrorMsg = "internal server error"
)

var (
	l                        = log.New(os.Stdout, ">> ", 16)
	rdb RedisClientInterface = &redis.Client{}
)

func main() {
	l.Println("starting visits app")

	l.Println("initialising Redis")
	redisHost := os.Getenv(redisHostNameKey)
	if redisHost == "" {
		l.Printf("%s env var is not set, using localhost\n", redisHostNameKey)
		redisHost = "localhost"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", redisHost),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	l.Println("pinging Redis")
	if err := rdb.Ping().Err(); err != nil {
		l.Println("redis ping returned an error", err)
	}

	http.HandleFunc("/visits", visitsEndpoint)

	l.Println("starting the server on port: 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		l.Fatal("server returned an error", err)
	}

	l.Println("server has stopped")
}

func visitsEndpoint(w http.ResponseWriter, r *http.Request) {
	l.Println("received a visit, processing visit counter")

	val, err := rdb.Get(redisVisitsKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			l.Printf("%q key doesn't exist in Redis yet", redisVisitsKey)

			val = "0"
		} else {
			l.Println("failed to get data from Redis", err)

			_, _ = w.Write([]byte(fmt.Sprintf("%s: %s", internalServerErrorMsg, "redis failed to respond")))
			return
		}
	}

	visits, err := strconv.Atoi(val)
	if err != nil {
		l.Println("failed to convert visits strings to int", err)

		_, _ = w.Write([]byte(internalServerErrorMsg))
		return
	}

	_, err = rdb.Set(redisVisitsKey, visits+1, 0).Result()
	if err != nil {
		l.Println("failed to set data in Redis")

		_, _ = w.Write([]byte(internalServerErrorMsg))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("number of visits: %d", visits+1)))
}
