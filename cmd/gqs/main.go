package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: check if static files are present if not maybe regenerate?
	// TODO: dont use the defaults

	prod := flag.Bool("prod", false, "data source from which to generate templates")

	flag.Parse()

	r := gin.Default()
	r.LoadHTMLFiles("assets/static/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Static("/assets", "./assets/dist")

	if *prod {
		r.RunTLS(":443", "/etc/letsencrypt/live/graphquasar.com/fullchain.pem", "/etc/letsencrypt/live/graphquasar.com/privkey.pem")
	} else {
		r.Run(":8080")
	}

}
