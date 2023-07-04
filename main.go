package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	ok := true

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	g := gin.Default()

	hn, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	g.LoadHTMLFiles("./static/index.html")
	g.GET("/", func(ctx *gin.Context) {
		statusstr := "NG"
		if ok {
			statusstr = "OK"
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{"HOSTNAME": hn, "STATUS": statusstr})
	})

	g.GET("/healthz", func(ctx *gin.Context) {
		if ok {
			ctx.String(http.StatusOK, "OK")
		} else {
			ctx.String(http.StatusInternalServerError, "NG")
		}
	})

	g.GET("/secret", func(ctx *gin.Context) {
		secret := os.Getenv("MY_SECRET")
		text := "Oh. I don't know your secrets!"
		if secret != "" {
			text = "Psst. Your secret is " + secret + "!"
		}
		ctx.String(http.StatusOK, text)
	})
	g.GET("/config", func(ctx *gin.Context) {
		config := os.Getenv("MY_CONFIG")
		text := "The server is not configured yet!"
		if config != "" {
			text = "OK! The server is configured as " + config + "!"
		}
		ctx.String(http.StatusOK, text)
	})

	g.GET("/butterfly", func(ctx *gin.Context) {
		ok = false
		fmt.Fprintln(os.Stderr, "ERROR: BUG: Haha! Nobody find... Oh no! You caught me!")
		ctx.String(http.StatusInternalServerError, "Something goes wrong")
	})

	g.StaticFile("/secretfile", "/app/secret/secretfile")
	g.StaticFile("/configfile", "/app/config/configfile")

	g.Run(":" + port)
}
