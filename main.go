package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	router := gin.Default()
	mel := melody.New()

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.GET("/ws", func(ctx *gin.Context) {
		if err := mel.HandleRequest(ctx.Writer, ctx.Request); err != nil {
			fmt.Println(err)
		}
	})

	mel.HandleMessage(func(session *melody.Session, msg []byte) {
		if err := mel.Broadcast(msg); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s -> %s \n", time.Now().Format(time.ANSIC), (msg))
	})

	if err := router.Run(":8010"); err != nil {
		panic(err)
	}
}
