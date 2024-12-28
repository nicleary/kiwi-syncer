package zims

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kiwi-syncer/external_services/kiwix_library"
	"net/http"
)

func GetZims(c *gin.Context) {
	var feed *kiwix_library.Feed
	var err error

	feed, err = kiwix_library.GetAvailableZims(0, 1000, "eng")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, feed)
}

func SubscribeToZim(c *gin.Context) {
	var ZimName ZimName
	if err := c.ShouldBindUri(&ZimName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid zim name or zim name not provided",
		})
		return
	}

	exists, err := kiwix_library.ZimExistsByName(ZimName.ZimName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error checking Zim status",
		})
		return
	}

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid zim name: Zim library does not any zim with the name %s", ZimName.ZimName),
		})
		return
	}

	//redisClient := redis.GetRedisClient()

	//redisClient.Set(c, "thingykey", "something", 0)
	//redisClient.HSet(c, "mysupermap", models.ZimSubscription{Name: "test", ID: "test", LastUpdatedAt: time.Now()})
}

func Routes(router *gin.Engine) {
	message := router.Group("api/v1/zim")
	{
		message.GET("/", GetZims)
		message.POST("/:id", SubscribeToZim)
	}
}
