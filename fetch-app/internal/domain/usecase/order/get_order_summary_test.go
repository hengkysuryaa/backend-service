package order

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/dto"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/entity"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/repository"
	webRepoMocks "github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/repository/mocks"
	"github.com/hengkysuryaa/backend-service/fetch-app/pkg/cache"
)

func Test_order_GetSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWebRepo := webRepoMocks.NewMockWebRepo(ctrl)
	startDate, _ := time.Parse("2006-01-02", "2022-01-06")
	cutoffDate := startDate.AddDate(0, 0, 7)

	type fields struct {
		webRepo repository.WebRepo
		cache   *cache.MapCache
	}
	type args struct {
		ctx    context.Context
		filter dto.GetSummaryFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.OrderSummary
		wantErr bool
		before  func()
	}{
		{
			name: "success get summary",
			fields: fields{
				webRepo: mockWebRepo,
			},
			args: args{
				ctx: context.Background(),
				filter: dto.GetSummaryFilter{
					AreaProvinsi: "jawa barat",
					Tanggal:      "2022-01-06",
				},
			},
			want: dto.OrderSummary{
				AreaProvinsi: "jawa barat",
				StartDate:    startDate,
				EndDate:      cutoffDate,
				Price: dto.Stats{
					Min:     1000,
					Max:     2000,
					Median:  1500,
					Average: 1500,
					Sum:     3000,
					List:    []float64{1000, 2000},
				},
				Size: dto.Stats{
					Min:     10,
					Max:     20,
					Median:  15,
					Average: 15,
					Sum:     30,
					List:    []float64{10, 20},
				},
			},
			wantErr: false,
			before: func() {
				uuid := []string{"abcd-efgh", "wxyz-hjka"}
				komoditas := []string{"nila", "bandeng"}
				areaProvinsi := "jawa barat"
				areaKota := "bandung"
				size := []string{"10", "20"}
				price := []string{"1000", "2000"}
				tglParsed := []string{"2022-01-11T18:59:42Z", "2022-01-12T01:52:15Z"}
				timestamp := []string{"1641927582674", "1641952335612"}
				mockWebRepo.EXPECT().GetOrders(gomock.Any()).Return([]entity.Order{
					{
						UUID:         &uuid[0],
						Komoditas:    &komoditas[0],
						AreaProvinsi: &areaProvinsi,
						AreaKota:     &areaKota,
						Size:         &size[0],
						Price:        &price[0],
						TglParsed:    &tglParsed[0],
						Timestamp:    &timestamp[0],
					},
					{
						UUID:         &uuid[1],
						Komoditas:    &komoditas[1],
						AreaProvinsi: &areaProvinsi,
						AreaKota:     &areaKota,
						Size:         &size[1],
						Price:        &price[1],
						TglParsed:    &tglParsed[1],
						Timestamp:    &timestamp[1],
					},
				}, nil)
			},
		},
		{
			name: "negative case: invalid request date format",
			fields: fields{
				webRepo: mockWebRepo,
			},
			args: args{
				ctx: context.Background(),
				filter: dto.GetSummaryFilter{
					AreaProvinsi: "jawa barat",
					Tanggal:      "2022/01/06",
				},
			},
			want:    dto.OrderSummary{},
			wantErr: true,
			before: func() {

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()
			u := &order{
				webRepo: tt.fields.webRepo,
				cache:   tt.fields.cache,
			}
			got, err := u.GetSummary(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("order.GetSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("order.GetSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
