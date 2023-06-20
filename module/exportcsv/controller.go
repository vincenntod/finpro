package exportcsv

import (
	"errors"
)

type ExportCSVController struct {
	useCase *UseCase
}

func NewController(useCase *UseCase) *ExportCSVController {
	return &ExportCSVController{
		useCase: useCase,
	}
}

type ExportCSVResponse struct {
	ExportCSVResponseMasage string `json:"masage"`
}

type TransactionCollection struct {
	DataItem []Transaction
}

func (c ExportCSVController) ExportCSV() ([][]string, error) {
	exportData, err := c.useCase.ExportCSV()
	if err != nil {
		return nil, err
	}
	return exportData, nil
}
func (c ExportCSVController) ExportCSVRangeDateFilter(startDate string, endDate string) ([][]string, error) {
	exportData, err := c.useCase.ExportCSVRangeDateFilter(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return exportData, nil
}
func (c ExportCSVController) ExportCSVStatusFilter(status string) ([][]string, error) {
	switch status {
	case "SUCCESS", "WAITING_FOR_DEBITTED":
		exportData, err := c.useCase.ExportCSVStatusFilter(status)
		if err != nil {
			return nil, err
		}
		return exportData, nil

	default:
		//kurang set error
		return nil, errors.New("Invalid field status")
	}

}
func (c ExportCSVController) ExportCSVStatusAndRangeDate(status string, startDate string, endDate string) ([][]string, error) {
	switch status {
	case "SUCCESS", "WAITING_FOR_DEBITTED":
		exportData, err := c.useCase.ExportCSVStatusAndRangeDate(status, startDate, endDate)
		if err != nil {
			return nil, err
		}
		return exportData, nil

	default:
		//kurang set error
		return nil, errors.New("Invalid field status")
	}

}
