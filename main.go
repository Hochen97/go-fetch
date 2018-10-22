package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"unsafe"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	// router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.LoadHTMLFiles("views/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	api.GET("/listImages", listImagesHandler)
	api.POST("/addImage", addImageHandler)
	api.POST("/delImage", delImageHandler)

	// Start and run the server
	router.Run(":3000")
}

// Utilities
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

// Websocket stuff
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Filter struct {
	Track string
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		msg1 := BytesToString(msg)
		fmt.Println(msg1)
		var filter Filter
		json.Unmarshal([]byte(msg), &filter)

		if err != nil {
			break
		}

		demux := twitter.NewSwitchDemux()
		demux.Tweet = func(tweet *twitter.Tweet) {
			// fmt.Println(tweet.ExtendedEntities.Media)
			for index, media := range tweet.ExtendedEntities.Media {
				fmt.Println(index, media.MediaURLHttps)
			}
		}

		config := oauth1.NewConfig("5CBnzw9q88yCCYjcBKvjpLpY5", "gNhqhZDh2pqIUULrRqOhPJ9lkJJdTr0sr8u12jn8I2DjEI0vKv")
		token := oauth1.NewToken("78039422-qykWx0nycETVmFj9hFT046fRid8BNKwgwFpmoIvlo", "UqWPxJsiLbViYAFwFBXKKkRHHGGXc4Snhhze7aZicmWdS")
		httpClient := config.Client(oauth1.NoContext, token)
		client := twitter.NewClient(httpClient)
		params := &twitter.StreamFilterParams{
			Track:         []string{filter.Track},
			StallWarnings: twitter.Bool(true),
		}
		stream, err := client.Streams.Filter(params)
		if err != nil {
			fmt.Println("Errpr in stream")
		}
		fmt.Println("Stream started with filter: %s", filter.Track)

		go demux.HandleChan(stream.Messages)

		conn.WriteMessage(t, msg)
	}
}

// Image Handling
func listImagesHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "listImages handler not implemented yet",
	})
}

func addImageHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "image addition handler not implemented yet",
	})
}

func delImageHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "image deletion handler not implemented yet",
	})
}

func imageListGenerator(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "list generator not implemented yet",
	})
}
