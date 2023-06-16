package exportcsv

import (
	"errors"
	"fmt"
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
	fmt.Println(req.Status)
	switch req.Status {
	case "SUCCESS", "WAITING_FOR_DEBITTED", "":
		exportData, err := c.useCase.ExportCSV(req)
		if err != nil {
			return nil, err
		}
		return exportData, nil

	default:
		//kurang set error
		return nil, errors.New("invalid field status")
	}

}
