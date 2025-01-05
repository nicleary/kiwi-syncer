package zims

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kiwi-syncer/external_services/kiwix_library"
	"kiwi-syncer/models"
	"kiwi-syncer/redis"
	"net/http"
	"time"
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

	redisClient := redis.GetRedisClient()

	// If zim is already subscribed to, do not go further
	if redisClient.Exists(c, ZimName.ZimName).Val() != 0 {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"subscribed": fmt.Sprintf("%s is already subscribed", ZimName.ZimName),
			})
		return
	}

	var subscriptionObject models.ZimSubscription

	// Get object from redis
	if err := redisClient.HGetAll(c, ZimName.ZimName).Scan(&subscriptionObject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error checking Zim status",
		})
		return
	}

	zim, err := kiwix_library.GetZimByName(ZimName.ZimName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error checking Zim status",
		})
		return
	}

	if zim == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid zim name: Zim library does not any zim with the name %s", ZimName.ZimName),
		})
		return
	}

	redisClient.HSet(c, ZimName.ZimName, models.ZimSubscription{SubscribedAt: time.Now(), LastUpdatedAt: time.Now(), LastSyncedAt: time.Now(), Name: zim.Name, Summary: zim.Summary, CurrentID: zim.Id.String()})

	c.JSON(http.StatusOK,
		gin.H{
			"Success": fmt.Sprintf("Successfully subscribed to zim %s", ZimName.ZimName),
		})
}

func Routes(router *gin.Engine) {
	message := router.Group("api/v1/zim")
	{
		message.GET("/", GetZims)
		message.POST("/:id", SubscribeToZim)
	}
}
