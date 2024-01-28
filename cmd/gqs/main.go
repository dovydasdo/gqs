package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: check if static files are present if not maybe regenerate?
	// TODO: dont use the defaults
	r := gin.Default()
	r.LoadHTMLFiles("assets/static/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Static("/assets", "./assets/dist")

	http.ListenAndServe(":8080", r)
}
