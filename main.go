package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	"kiwi-syncer/apis/zims"
	"kiwi-syncer/redis"
	"log"
	"net/http"
)

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
	r.Run()
}
