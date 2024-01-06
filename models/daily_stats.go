package models

import "time"

type DailyStatsByCity struct {
	AveragePrice          int
	AveragePricePerSquare float64
	AverageFootage        float64
	AdsCount              int
	Date                  time.Time
	City                  string
}
