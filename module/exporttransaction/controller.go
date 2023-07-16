package exporttransaction

import (
	"errors"
	"fmt"
	"golang/module/exporttransaction/dto"
	"golang/module/exporttransaction/entity"
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
	DataItem []entity.Transaction
}
type Controller interface {
	ExportCSV(req *dto.ExportCSVRequest) ([][]string, error)
}

func (c controller) ExportCSV(req *dto.ExportCSVRequest) ([][]string, error) {
	isRequestInvalid := false
	var clientWant = req.Url.ClientUrl

	var serverGot string
	fmt.Println(req)
	switch {
	case len(req.Filter.StartDate) != 10 && req.Filter.StartDate != "" || len(req.Filter.EndDate) != 10 && req.Filter.EndDate != "":
		isRequestInvalid = true
	case (req.Filter.StartDate == "" && req.Filter.EndDate != "") || (req.Filter.StartDate != "" && req.Filter.EndDate == ""):
		isRequestInvalid = true
	case req.Filter.StartDate != "" && req.Filter.StartDate == "" && req.Filter.EndDate == "":
		serverGot = req.Url.PathUrl + "?" + "status=" + req.Filter.Status
		if clientWant != serverGot {
			isRequestInvalid = true
		}
	case req.Filter.Status == "" && req.Filter.StartDate != "" && req.Filter.EndDate != "":
		serverGot = req.Url.PathUrl + "?" + "start_date=" + req.Filter.StartDate + "&end_date=" + req.Filter.EndDate
		if serverGot != clientWant {
			isRequestInvalid = true
		}
	case req.Filter.Status != "" && req.Filter.StartDate != "" && req.Filter.EndDate != "":
		serverGot = req.Url.PathUrl + "?" + "status=" + req.Filter.Status + "&start_date=" + req.Filter.StartDate + "&end_date=" + req.Filter.EndDate
		if serverGot != clientWant {
			isRequestInvalid = true
		}
	}

	if !isRequestInvalid {
		for _, sts := range entity.StatusArray {
			if sts == req.Filter.Status {
				exportData, err := c.useCase.ExportCSV(req)
				if err != nil {
					return nil, err
				}
				return exportData, nil
			}
		}
	}
	return nil, errors.New("invalid request")
}
