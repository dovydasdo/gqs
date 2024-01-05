package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dovydasdo/gqs/config"
	"github.com/dovydasdo/gqs/models"
	"github.com/jackc/pgx"
)

type PSQLDB struct {
	conn *pgx.Conn
}

func GetPSQLDB() (*PSQLDB, error) {
	conf := config.GetDBConfig()

	srcDBConnStr := fmt.Sprintf("postgresql://%v:%v@%v:5432/%v", conf.Username, conf.Password, conf.Host, conf.DBName)
	log.Println(srcDBConnStr)
	conConf, err := pgx.ParseConnectionString(srcDBConnStr)
	if err != nil {
		return nil, err
	}
	conn, err := pgx.Connect(conConf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &PSQLDB{
		conn: conn,
	}, nil
}

func (db *PSQLDB) GetDailyStatsByCity() ([]models.DailyStatsByCity, error) {

	rows, err := db.conn.Query("SELECT average_price, average_price_per_sq, average_footage, city, date, ads_count FROM rent_daily_city_stats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []models.DailyStatsByCity
	counter := 0

	for rows.Next() {
		var stat models.DailyStatsByCity

		if err := rows.Scan(&stat.AveragePrice, &stat.AveragePricePerSquare, &stat.AverageFootage, &stat.City, &stat.Date, &stat.AdsCount); err != nil {
			log.Fatal(err)
		}

		noTime := time.Time{}

		if stat.AveragePrice == 0 || stat.AveragePricePerSquare == 0 || stat.AverageFootage == 0 || stat.City == "" || stat.Date == noTime {
			continue
		}

		stats = append(stats, stat)
		counter++

	}

	return stats, nil
}
