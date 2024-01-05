package main

import (
	"context"
	"net/http"

	"github.com/dovydasdo/gqs/templates"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.LoadHTMLFiles("dist/index.html")
	r.GET("/", func(c *gin.Context) {
		comp := templates.Hello("beans")
		comp.Render(context.Background(), c.Writer)

	})

	http.ListenAndServe(":8080", r)
}
