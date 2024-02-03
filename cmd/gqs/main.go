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
	r.LoadHTMLFiles("assets/static/index.html", "assets/static/rent.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/rent", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rent.html", gin.H{})
	})

	r.Static("/assets", "./assets/dist")
	r.Static("/svgs", "./assets/svgs")

	if *prod {
		r.RunTLS(":443", "./fullchain.pem", "./privkey.pem")
	} else {
		r.Run(":8080")
	}

}
