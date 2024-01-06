package db

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dovydasdo/gqs/models"
)

// TODO: glob for caching
type LocalCSVCache struct {
	pathsToCache []string
}

func GetCSVCache(paths ...string) (LocalCSVCache, error) {
	lc := LocalCSVCache{}
	lc.pathsToCache = make([]string, 0)

	for _, path := range paths {
		lc.pathsToCache = append(lc.pathsToCache, path)
	}

	return lc, nil
}

func (lc LocalCSVCache) GetDailyStatsByCity() ([]models.DailyStatsByCity, error) {
	var data []models.DailyStatsByCity

	// todo: maybe table name should be configurable in some way?
	path, err := getFromArr(lc.pathsToCache, "rent_daily_city_stats")
	if err != nil {
		return nil, errors.New("no path for table found")
	}

	file, err := os.Open(path)
	if err != nil {
		return data, err
	}

	defer file.Close()

	csvr := csv.NewReader(file)

	for {
		row, err := csvr.Read()
		if err != nil {
			break
		}

		stat := models.DailyStatsByCity{}

		var price int64

		if price, err = strconv.ParseInt(row[4], 10, 64); err != nil {
			log.Printf("failed to parse price : val = %v, err: %v", row[4], err)
			continue
		}

		stat.AveragePrice = int(price)

		if stat.AveragePricePerSquare, err = strconv.ParseFloat(row[5], 64); err != nil {
			log.Printf("failed to parse price per square : val = %v, err: %v", row[5], err)
			continue
		}

		if stat.AverageFootage, err = strconv.ParseFloat(row[6], 64); err != nil {
			log.Printf("failed to parse footage : val = %v, err: %v", row[6], err)
			continue
		}

		if row[7] == "" {
			log.Printf("city was empty")
			continue
		}

		stat.City = row[7]

		if row[8] == "" {
			log.Printf("date was empty")
			continue
		}

		if stat.Date, err = time.Parse("2006-01-02", row[8]); err != nil {
			log.Printf("failed to parse footage : val = %v, err: %v", row[6], err)
			continue
		}

		var count int64

		if count, err = strconv.ParseInt(row[9], 10, 64); err != nil {
			log.Printf("failed to parse count: val = %v, err: %v", row[9], err)
			continue
		}

		stat.AdsCount = int(count)

		data = append(data, stat)
	}

	return data, nil
}

func getFromArr(source []string, target string) (string, error) {
	for _, s := range source {
		if strings.Contains(s, target) {
			return s, nil
		}
	}

	return "", errors.New("string not found")
}
