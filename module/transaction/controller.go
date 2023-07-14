package transaction

import (
	"golang/module/transaction/dto"
	"golang/module/transaction/entities"
	"golang/module/transaction/helper"
	"net/http"
)

type Controller struct {
	UseCase UseCaseInterface
}

type ControllerInterface interface {
	GetAllTransaction(page int, size int) (*dto.GetAllResponseDataTransaction, error)
	GetAllTransactionByStatus(status string, page int, size int) (*dto.GetAllResponseDataTransaction, error)
	GetAllTransactionByDate(start string, end string, page int, size int) (*dto.GetAllResponseDataTransaction, error)
	GetAllTransactionByStatusDate(status string, start string, end string, page int, size int) (*dto.GetAllResponseDataTransaction, error)

	GetAllTransactionNoLimit() (*dto.GetAllResponseDataTransaction, error)
	GetAllTransactionByStatusNoLimit(status string) (*dto.GetAllResponseDataTransaction, error)
	GetAllTransactionByDateNoLimit(start string, end string) (*dto.GetAllResponseDataTransaction, error)
	GetAllTransactionByStatusDateNoLimit(status string, start string, end string) (*dto.GetAllResponseDataTransaction, error)

	GetAllTransactionByRequest(req *dto.Request) (*dto.GetAllResponseDataTransaction, error)
}

func NewController(useCase UseCaseInterface) ControllerInterface {
	return Controller{
		UseCase: useCase,
	}
}

func (Controller) DataResponse(transaction []entities.Transaction) (*dto.GetAllResponseDataTransaction, error) {

	if len(transaction) == 0 {
		res := &dto.GetAllResponseDataTransaction{
			Code:      http.StatusBadRequest,
			Message:   "Data Tidak Ditemukan Atau Inputan Salah",
			Error:     "Not Success",
			TotalData: len(transaction),
		}

		return res, nil
	} else {

		res := &dto.GetAllResponseDataTransaction{
			Code:      http.StatusOK,
			Message:   "Data Berhasil Diambil",
			Error:     "Success",
			TotalData: len(transaction),
		}

		for _, v := range transaction {
			res.Data = append(res.Data, dto.TransactionItemResponse{
				Id:         v.Id,
				OdaNumber:  v.OdaNumber,
				BillAmount: v.BillAmount,
				Status:     v.Status,
				CreatedAt:  v.CreatedAt,
			})
		}
		return res, nil
	}
}

func (Controller) InputFormatDate(req *dto.Request) (string, string) {
	start := helper.FormatDate(req.StartDate)
	end := helper.FormatDate(req.EndDate)
	return start, end
}

func (c Controller) GetAllTransaction(page int, size int) (*dto.GetAllResponseDataTransaction, error) {

	transaction, err := c.UseCase.GetAllTransaction(page, size)
	if err != nil {
		return nil, err
	}

	return c.DataResponse(transaction)
}

func (c Controller) GetAllTransactionByStatus(status string, page int, size int) (*dto.GetAllResponseDataTransaction, error) {

	transaction, err := c.UseCase.GetAllTransactionByStatus(status, page, size)
	if err != nil {
		return nil, err
	}

	return c.DataResponse(transaction)
}

func (c Controller) GetAllTransactionByDate(start string, end string, page int, size int) (*dto.GetAllResponseDataTransaction, error) {
	transaction, err := c.UseCase.GetAllTransactionByDate(start, end, page, size)
	if err != nil {
		res := &dto.GetAllResponseDataTransaction{
			Code:      http.StatusBadRequest,
			Message:   "Bad Request",
			Error:     "Error",
			TotalData: len(transaction),
		}
		return res, nil
	}

	return c.DataResponse(transaction)
}

func (c Controller) GetAllTransactionByStatusDate(status string, start string, end string, page int, size int) (*dto.GetAllResponseDataTransaction, error) {
	transaction, err := c.UseCase.GetAllTransactionByStatusDate(status, start, end, page, size)
	if err != nil {
		return nil, err
	}

	return c.DataResponse(transaction)
}

func (c Controller) GetAllTransactionNoLimit() (*dto.GetAllResponseDataTransaction, error) {

	transaction, err := c.UseCase.GetAllTransactionNoLimit()
	if err != nil {
		return nil, err
	}

	return c.DataResponse(transaction)
}

func (c Controller) GetAllTransactionByStatusNoLimit(status string) (*dto.GetAllResponseDataTransaction, error) {

	transaction, err := c.UseCase.GetAllTransactionByStatusNoLimit(status)
	if err != nil {
		return nil, err
	}

	return c.DataResponse(transaction)
}

func (c Controller) GetAllTransactionByDateNoLimit(start string, end string) (*dto.GetAllResponseDataTransaction, error) {
	transaction, err := c.UseCase.GetAllTransactionByDateNoLimit(start, end)
	if err != nil {
		res := &dto.GetAllResponseDataTransaction{
			Code:      http.StatusBadRequest,
			Message:   "Bad Request",
			Error:     "Error",
			TotalData: len(transaction),
		}
		return res, nil
	}

	return c.DataResponse(transaction)
}

func (c Controller) GetAllTransactionByStatusDateNoLimit(status string, start string, end string) (*dto.GetAllResponseDataTransaction, error) {
	transaction, err := c.UseCase.GetAllTransactionByStatusDateNoLimit(status, start, end)
	if err != nil {
		return nil, err
	}

	return c.DataResponse(transaction)
}

// Method kontroler untuk mengambil data transaksi berdasarkan request

func (c Controller) GetAllTransactionByRequest(req *dto.Request) (*dto.GetAllResponseDataTransaction, error) {

	res := &dto.GetAllResponseDataTransaction{
		Code:    http.StatusBadRequest,
		Message: "Inputan Salah",
		Error:   "Bad Request",
		Data:    nil,
	}

	switch {
	case req.Status == "" && req.StartDate == "" && req.EndDate == "" && req.Page != 0 && req.Size != 0:

		return c.GetAllTransaction(req.Page, req.Size)

	case req.Status != "" && req.StartDate == "" && req.EndDate == "" && req.Page != 0 && req.Size != 0:

		return c.GetAllTransactionByStatus(req.Status, req.Page, req.Size)

	case req.Status == "" && req.StartDate != "" && req.EndDate != "" && req.Page != 0 && req.Size != 0:

		start, end := c.InputFormatDate(req)
		return c.GetAllTransactionByDate(start, end, req.Page, req.Size)

	case req.Status != "" && req.StartDate != "" && req.EndDate != "" && req.Page != 0 && req.Size != 0:

		start, end := c.InputFormatDate(req)
		return c.GetAllTransactionByStatusDate(req.Status, start, end, req.Page, req.Size)

	case req.Status == "" && req.StartDate == "" && req.EndDate == "" && req.Page == 0 && req.Size == 0:

		return c.GetAllTransactionNoLimit()

	case req.Status != "" && req.StartDate == "" && req.EndDate == "" && req.Page == 0 && req.Size == 0:

		return c.GetAllTransactionByStatusNoLimit(req.Status)

	case req.Status == "" && req.StartDate != "" && req.EndDate != "" && req.Page == 0 && req.Size == 0:

		start, end := c.InputFormatDate(req)
		return c.GetAllTransactionByDateNoLimit(start, end)

	case req.Status != "" && req.StartDate != "" && req.EndDate != "" && req.Page == 0 && req.Size == 0:

		start, end := c.InputFormatDate(req)
		return c.GetAllTransactionByStatusDateNoLimit(req.Status, start, end)

	default:
		return res, nil
	}
}
