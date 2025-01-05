package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/onatm/clockwerk"
	stats "github.com/semihalev/gin-stats"
	"kiwi-syncer/apis/zims"
	"kiwi-syncer/redis"
	"log"
	"net/http"
	"time"
)

type DummyJob struct{}

func (d DummyJob) Run() {
	fmt.Println("Every 30 seconds")
}

func main() {
	fmt.Println("Starting server")

	err := redis.InitializeRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(stats.RequestStats())

	r.GET("api/v1/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	zims.Routes(r)
	var job DummyJob
	c := clockwerk.New()
	c.Every(30 * time.Second).Do(job)
	//c.Start()
	r.Run()
}
