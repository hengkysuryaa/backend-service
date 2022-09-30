package request

type GetSummaryRequest struct {
	AreaProvinsi string `json:"area_provinsi" validate:"required"`
	Tanggal      string `json:"tanggal" validate:"required"`
}
