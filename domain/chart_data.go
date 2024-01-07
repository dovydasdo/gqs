package domain

import (
	"errors"
	"sort"
	"time"

	"github.com/dovydasdo/gqs/models"
)

func GetStatsByCity(src []models.DailyStatsByCity) ([]models.StatsByCity, error) {
	if len(src) == 0 {
		return nil, errors.New("not data was provided")
	}

	cities := getCities(src)
	statsByCity := make([]models.StatsByCity, 0)
	allStats := src

	for city := range cities {

		stats := models.StatsByCity{}
		stats.City = city

		dc, reduced := getWithString(city, allStats)
		allStats = reduced

		// sanity check
		if len(dc) < 1 {
			break
		}

		prices := make(map[time.Time]int, 0)
		listings := make(map[time.Time]int, 0)
		pps := make(map[time.Time]float64, 0)
		footage := make(map[time.Time]float64, 0)

		for _, s := range dc {
			prices[s.Date] = s.AveragePrice
			listings[s.Date] = s.AdsCount
			pps[s.Date] = s.AveragePricePerSquare
			footage[s.Date] = s.AverageFootage
		}

		// //sort by date is not possible for maps
		// prices = sortByDate[int, map[time.Time]int](prices)
		// listings = sortByDate[int, map[time.Time]int](listings)
		// pps = sortByDate[float64, map[time.Time]float64](pps)
		// footage = sortByDate[float64, map[time.Time]float64](footage)

		stats.CountStats = listings
		stats.PriceStats = prices
		stats.PPSStats = pps
		stats.FootageStats = footage
		stats.AvgListings = getSum[time.Time, int](listings) / len(listings)
		stats.SortedDates = getSortedDates[int](prices)

		statsByCity = append(statsByCity, stats)

		// no more stats so we can stop
		if len(allStats) < 1 {
			break
		}

	}

	sort.Slice(statsByCity, func(i, j int) bool {
		return statsByCity[i].AvgListings > statsByCity[j].AvgListings
	})

	return statsByCity, nil
}

func getSum[K time.Time, V int | float64](v map[K]V) V {
	var sum V
	for _, d := range v {
		sum += d
	}

	return sum
}

func getSortedDates[V int | float64](v map[time.Time]V) []time.Time {
	keys := make([]time.Time, 0)

	for t := range v {
		keys = append(keys, t)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	return keys
}

func getCities(d []models.DailyStatsByCity) map[string]struct{} {
	r := make(map[string]struct{}, 0)
	for _, c := range d {
		if _, ok := r[c.City]; ok {
			continue
		}

		r[c.City] = struct{}{}
	}

	return r
}

// plz test this shit
func getWithString(s string, d []models.DailyStatsByCity) ([]models.DailyStatsByCity, []models.DailyStatsByCity) {
	// god forgive me for what im about to do
	found := 0
	// put elements with the desired string in the front
	for i, c := range d {
		if c.City == s {
			tmp := d[found]
			d[found] = c
			d[i] = tmp

			found++
		}
	}

	// from start to found is with string, others are without
	return d[:found], d[found:]
}
