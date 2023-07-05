package exportcsv

import (
	"errors"
)

type controller struct {
	useCase Usecase
}

func NewController(useCase Usecase) *controller {
	return &controller{
		useCase: useCase,
	}
}

type ExportCSVResponse struct {
	ExportCSVResponseMasage string `json:"masage"`
}

type TransactionCollection struct {
	DataItem []Transaction
}
type Controller interface {
	ExportCSV(req *ExportCSVRequest) ([][]string, error)
}

func (c controller) ExportCSV(req *ExportCSVRequest) ([][]string, error) {
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
