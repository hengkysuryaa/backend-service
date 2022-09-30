package dto

import "time"

type Stats struct {
	Min     float64   `json:"min"`
	Max     float64   `json:"max"`
	Median  float64   `json:"median"`
	Average float64   `json:"average"`
	Sum     float64   `json:"-"`
	List    []float64 `json:"-"`
}

type GetSummaryFilter struct {
	AreaProvinsi string
	Tanggal      string
}

type OrderSummary struct {
	AreaProvinsi string    `json:"area_provinsi"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Price        Stats     `json:"price"`
	Size         Stats     `json:"size"`
}
