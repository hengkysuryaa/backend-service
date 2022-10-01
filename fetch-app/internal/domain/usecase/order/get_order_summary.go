package order

import (
	"context"
	"log"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/pkg/helpers"
)

func (u *order) GetSummary(ctx context.Context, filter dto.GetSummaryFilter) (dto.OrderSummary, error) {
	var priceStats dto.Stats
	var sizeStats dto.Stats

	startDate, err := time.Parse("2006-01-02", filter.Tanggal)
	if err != nil {
		log.Println(err)
		return dto.OrderSummary{}, err
	}

	// weekly summary, add filter date by 7 day
	cutOffDate := startDate.AddDate(0, 0, 7)

	orders, err := u.webRepo.GetOrders(ctx)
	if err != nil {
		log.Println(err)
		return dto.OrderSummary{}, err
	}

	for i, order := range orders {
		// remove data outlier
		if order.AreaProvinsi == nil || order.TglParsed == nil || order.Timestamp == nil {
			continue
		}

		if order.Price == nil || order.Size == nil {
			continue
		}

		// seed first data
		if i == 0 {
			price, _ := strconv.ParseFloat(*order.Price, 64)
			size, _ := strconv.ParseFloat(*order.Size, 64)
			priceStats.Min = price
			sizeStats.Min = size
		}

		// by area_provinsi and weekly range
		if helpers.DeReferenceString(order.AreaProvinsi) == filter.AreaProvinsi {
			timestamp, _ := strconv.ParseInt(*order.Timestamp, 10, 64)
			startDateMilli := startDate.UnixNano() / int64(time.Millisecond)
			cutOffDateMilli := cutOffDate.UnixNano() / int64(time.Millisecond)

			if !(timestamp >= startDateMilli && timestamp <= cutOffDateMilli) {
				continue
			}

			u.getMinMaxStats(&priceStats, *order.Price)
			u.getMinMaxStats(&sizeStats, *order.Size)
		}
	}

	// get median and average
	priceStats.Median = u.getMedianStats(priceStats.List)
	priceStats.Average = u.getAverageStats(priceStats.Sum, float64(len(priceStats.List)))

	sizeStats.Median = u.getMedianStats(sizeStats.List)
	sizeStats.Average = u.getAverageStats(sizeStats.Sum, float64(len(sizeStats.List)))

	return dto.OrderSummary{
		AreaProvinsi: filter.AreaProvinsi,
		StartDate:    startDate,
		EndDate:      cutOffDate,
		Price:        priceStats,
		Size:         sizeStats,
	}, nil
}

func (u *order) getMinMaxStats(stats *dto.Stats, value string) {
	val, _ := strconv.ParseFloat(value, 64)

	stats.Min = math.Min(float64(stats.Min), val)
	stats.Max = math.Max(float64(stats.Max), val)
	stats.Sum += val
	stats.List = append(stats.List, val)
}

func (u *order) getMedianStats(list []float64) float64 {
	if len(list) == 0 {
		return 0
	}

	var median float64
	sort.Float64s(list)
	n := len(list)

	if n%2 == 0 {
		median = (list[n/2-1] + list[n/2]) / 2
	} else {
		median = list[n/2]
	}

	return median
}

func (u *order) getAverageStats(sum, len float64) float64 {
	if len == 0 {
		return 0
	}

	return sum / len
}
