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

type Order struct {
	UUID         *string  `json:"uuid"`
	Komoditas    *string  `json:"komoditas"`
	AreaProvinsi *string  `json:"area_provinsi"`
	AreaKota     *string  `json:"area_kota"`
	Size         *string  `json:"size"`
	IDRPrice     *float64 `json:"idr_price"`
	USDPrice     *float64 `json:"usd_price"`
	TglParsed    *string  `json:"tgl_parsed"`
	Timestamp    *string  `json:"timestamp"`
}
