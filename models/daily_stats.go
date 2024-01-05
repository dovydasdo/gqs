package models

import "time"

type DailyStatsByCity struct {
	AveragePrice          int
	AveragePricePerSquare float32
	AverageFootage        float32
	AdsCount              int
	Date                  time.Time
	City                  string
}
