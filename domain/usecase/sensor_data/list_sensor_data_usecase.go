package sensor_data

import (
	"context"
	"device-service/domain/entity"
	"device-service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListSensorDataRequest struct {
	Filters    SensorDataFilters
	Pagination *common.Pagination
}

type ListSensorDataResponse common.PaginationResult[SensorDataItem]

type SensorDataItem struct {
	ID           string
	DeviceID     string
	SensorType   string
	Value        float64
	Unit         string
	RecordedAt   time.Time
	IsAlert      bool
	QualityScore float64
	CreatedAt    time.Time
}

type SensorDataFilters struct {
	DeviceID   *string
	SensorType *string
	IsAlert    *bool
	StartDate  *string
	EndDate    *string
}

type ListSensorDataUsecase struct {
	sensorDataRepo repository.SensorDataRepository
	helper         utils.Helper
}

func NewListSensorDataUsecase(sensorDataRepo repository.SensorDataRepository, helper utils.Helper) *ListSensorDataUsecase {
	return &ListSensorDataUsecase{
		sensorDataRepo: sensorDataRepo,
		helper:         helper,
	}
}

func (u *ListSensorDataUsecase) Execute(ctx context.Context, req *ListSensorDataRequest) (*ListSensorDataResponse, error) {

	pagination := req.Pagination
	if pagination == nil {
		pagination = &common.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}

	filters := repository.SensorDataFilters{
		DeviceID: req.Filters.DeviceID,
		IsAlert:  req.Filters.IsAlert,
	}

	if req.Filters.SensorType != nil {
		sensorType := entity.SensorType(*req.Filters.SensorType)
		filters.SensorType = &sensorType
	}

	if req.Filters.StartDate != nil {
		startDate, err := time.Parse("2006-01-02", *req.Filters.StartDate)
		if err != nil {
			return nil, ErrInvalidDateFormat
		}
		filters.StartDate = &startDate
	}

	if req.Filters.EndDate != nil {
		endDate, err := time.Parse("2006-01-02", *req.Filters.EndDate)
		if err != nil {
			return nil, ErrInvalidDateFormat
		}
		filters.EndDate = &endDate
	}

	sensorDataList, total, err := u.sensorDataRepo.List(ctx, filters, pagination)
	if err != nil {
		return nil, err
	}

	items := make([]SensorDataItem, len(sensorDataList))
	for i, sensorData := range sensorDataList {
		item := SensorDataItem{
			ID:           sensorData.ID,
			DeviceID:     sensorData.DeviceID,
			SensorType:   sensorData.SensorType,
			Value:        sensorData.Value,
			Unit:         sensorData.Unit,
			RecordedAt:   sensorData.RecordedAt,
			IsAlert:      sensorData.IsAlert,
			QualityScore: sensorData.QualityScore,
			CreatedAt:    sensorData.CreatedAt,
		}

		items[i] = item
	}

	return &ListSensorDataResponse{
		Data:       items,
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: u.helper.CalculateTotalPages(int64(total), int64(pagination.PageSize)),
	}, nil
}
