package generators

import (
	"context"
	"os"

	"github.com/dovydasdo/gqs/models"
	"github.com/dovydasdo/gqs/templates"
)

type MainStatsReader interface {
	GetDailyStatsByCity() ([]models.DailyStatsByCity, error)
}

// reads daily stats from db and generates html for index page
type MainGenerator struct {
	Reader MainStatsReader
}

func GetMainGenerator(reader MainStatsReader) *MainGenerator {
	return &MainGenerator{
		Reader: reader,
	}
}

func (g *MainGenerator) GenerateIndex(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	templates.Main().Render(context.Background(), file)

	return nil
}

func (g *MainGenerator) GenerateRentPage(path string) error {
	args, err := g.Reader.GetDailyStatsByCity()
	if err != nil {
		return err
	}

	// open file writer
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	templates.MainRentPage(args).Render(context.Background(), file)

	return nil
}
