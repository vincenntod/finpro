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

func (c ExportCSVController) ExportCSV(req *ExportCSVRequest) ([][]string, error) {
	switch req.Status {
	case "SUCCESS", "WAITING_FOR_DEBITTED", "": //validate request.status
		exportData, err := c.useCase.ExportCSV(req)
		if err != nil {
			return nil, err
		}
		return exportData, nil

	default:
		return nil, errors.New("Invalid field status")
	}

}
