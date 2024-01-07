// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package graphs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "log"
import "math"

import (
	"github.com/dovydasdo/gqs/domain"
	"github.com/dovydasdo/gqs/models"
	"github.com/dovydasdo/gqs/templates/shared"
)

func AllInfo(data []models.DailyStatsByCity) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			for _, c := range getStatsByCity(data) {
				templ_7745c5c3_Err = Graph(c, getMax(c), getMax(c)-getMin(c)).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = shared.Info("Rent graphs", "information about rent data...").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func getStatsByCity(d []models.DailyStatsByCity) []models.StatsByCity {
	stats, err := domain.GetStatsByCity(d)
	if err != nil {
		log.Printf("failed to parse data by city")
		return make([]models.StatsByCity, 0)
	}

	return stats
}

func getMax(d models.StatsByCity) int {
	max := math.MinInt

	for _, val := range d.PriceStats {
		if val > max {
			max = val
		}
	}

	if max == math.MinInt {
		log.Printf("failed to get min/max val for price")
		return 0
	}

	return max
}
func getMin(d models.StatsByCity) int {
	min := math.MaxInt

	for _, val := range d.PriceStats {
		if val < min {
			min = val
		}
	}

	if min == math.MinInt {
		log.Printf("failed to get min/max val for price")
		return 0
	}

	return min
}
