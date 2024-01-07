package models

import "time"

type StatsByCity struct {
	City         string
	AvgListings  int
	SortedDates  []time.Time
	PriceStats   map[time.Time]int
	CountStats   map[time.Time]int
	PPSStats     map[time.Time]float64
	FootageStats map[time.Time]float64
}
