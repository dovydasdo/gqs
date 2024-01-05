package generators

type MainStatsReader interface {
	GetDailyStatsByCity() ([]models.DailyStatsByCity, error)
}

type MainGenerator struct {
}
