package main

import (
	"flag"
	"log"

	"github.com/dovydasdo/gqs/db"
	"github.com/dovydasdo/gqs/generators"
)

func main() {
	source := flag.String("source", "psql", "data source from which to generate templates")

	flag.Parse()

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

	err = mainGen.GenerateIndex("./assets/static/index.html")
	if err != nil {
		log.Printf("Failed to generate files: %v", err)
	}

	err = mainGen.GenerateRentPage("./assets/static/rent.html")
	if err != nil {
		log.Printf("Failed to generate files: %v", err)
	}
}
