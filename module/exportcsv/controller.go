package exportcsv

import (
	"errors"
)

type exportCSVController struct {
	UseCase Usecase
}

func NewController(useCase Usecase) *exportCSVController {
	return &exportCSVController{
		UseCase: useCase,
	}
}

type ExportCSVResponse struct {
	ExportCSVResponseMasage string `json:"masage"`
}

type TransactionCollection struct {
	DataItem []Transaction
}
type ExportCSVController interface {
	ExportCSV(req *ExportCSVRequest) ([][]string, error)
}

func (c exportCSVController) ExportCSV(req *ExportCSVRequest) ([][]string, error) {
	switch req.Status {
	case "SUCCESS", "WAITING_FOR_DEBITTED", "": //validate request.status
		exportData, err := c.UseCase.ExportCSV(req)
		if err != nil {

			return nil, err
		}
		return exportData, nil

	default:
		return nil, errors.New("Invalid field status")
	}

}
