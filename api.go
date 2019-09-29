package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)

func main() {
    // Default returns an Engine instance with the Logger and Recovery middleware already attached
    r := gin.Default()

	channelItem := make(chan string)
	channelItem2 := make(chan string)

	bruh := func() {
		time.Sleep(time.Second * 1)
		channelItem <- "response de la api 1"
	}

	sis := func() {
		time.Sleep(time.Second * 3)
		channelItem2 <- "response de la api 2"
	}

	//GET of an single item full data
	r.GET("/full/:itemID", func(c *gin.Context){
		itemID := c.Param("itemID")
		go bruh()
		go sis()
		select{
		case primerParte := <-channelItem:
			fmt.Println("Primera parte:", primerParte)
		case segundaParte := <-channelItem2:
			fmt.Println("Segunda parte:", segundaParte)
		}
		time.Sleep(time.Second*10)
		c.JSON(200, gin.H{
			"asd": itemID,
		})
	})
    r.Run()
}