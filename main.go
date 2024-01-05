package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("dist/index.html")
	r.GET("/", func(c *gin.Context) {
		// c.Data(200, "text/html; charset=utf-8", []byte("cool"))
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	http.ListenAndServe(":8080", r)
}
