package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dovydasdo/gqs/db"
	"github.com/dovydasdo/gqs/generators"
	"github.com/gin-gonic/gin"
)

func main() {

	// var mode string

	mode := flag.String("mode", "none", "mode of running. SERVE for serving mode, GEN for regenerating templates")
	source := flag.String("source", "psql", "data source from which to generate templates")

	flag.Parse()

	// TODO this is not great, maybe seperating is a better idea
	switch *mode {
	case "serve":
		// TODO: check if static files are present if not maybe regenerate?

		// TODO: dont use the defaults
		r := gin.Default()
		r.LoadHTMLFiles("assets/static/index.html")
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})

		r.Static("/assets", "./assets/dist")

		http.ListenAndServe(":8080", r)
	case "gen":
		var reader generators.MainStatsReader
		var err error

		switch *source {
		case "psql":
			reader, err = db.GetPSQLDB()
			if err != nil {
				log.Panicf("connection to db failed: %v", err)
			}
		case "cache":
			reader, err = db.GetCSVCache("./.cache/rent_daily_city_stats_rows.csv")
			if err != nil {
				log.Panicf("failed to get data from local cache: %v", err)
			}
		default:
			log.Fatalf("the provided source is not valid")
		}

		mainGen := generators.GetMainGenerator(reader)

		err = mainGen.Generate("./assets/static/index.html")
		if err != nil {
			log.Printf("Failed to generate files: %v", err)
		}

	default:
		log.Fatalf("provided mode parameter is not valid: %v", *mode)
	}

}
