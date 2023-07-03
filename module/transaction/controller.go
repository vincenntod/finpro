package transaction

import (
	"net/http"
)

type Controller struct {
	useCase UseCaseInterface
}

type ControllerInterface interface {
	GetAllTransaction() (*GetAllResponseDataTransaction, error)
	GetAllTransactionByStatus(status string) (*GetAllResponseDataTransaction, error)
	GetAllTransactionByDate(start string, end string) (*GetAllResponseDataTransaction, error)
	GetAllTransactionByStatusDate(status string, start string, end string) (*GetAllResponseDataTransaction, error)
	GetTransaction(req *FilterByStatusDate) (*GetAllResponseDataTransaction, error)

	GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) (*GetAllResponseDataTransaction, error)
	GetTransactionByDate(req FilterByDate, input FilterLimit) (*GetAllResponseDataTransaction, error)
}

type TransactionItemResponse struct {
	Id         int     `json:"id"`
	OdaNumber  string  `json:"oda_number"`
	BillAmount float32 `json:"price"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"created_at"`
}

func NewController(useCase UseCaseInterface) ControllerInterface {
	return Controller{
		useCase: useCase,
	}
}

func (c Controller) GetAllTransaction() (*GetAllResponseDataTransaction, error) {

	transaction, err := c.useCase.GetAllTransaction()
	if err != nil {
		return nil, err
	}

	res := &GetAllResponseDataTransaction{
		Code:      200,
		Message:   "Data Berhasil Diambil",
		Error:     "Success",
		TotalData: len(transaction),
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:         v.Id,
			OdaNumber:  v.OdaNumber,
			BillAmount: v.BillAmount,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
		})
	}
	return res, nil
}

func (c Controller) GetAllTransactionByStatus(status string) (*GetAllResponseDataTransaction, error) {

	transaction, err := c.useCase.GetAllTransactionByStatus(status)
	if err != nil {
		return nil, err
	}

	res := &GetAllResponseDataTransaction{
		Code:      200,
		Message:   "Data Berhasil Diambil",
		Error:     "Success",
		TotalData: len(transaction),
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:         v.Id,
			OdaNumber:  v.OdaNumber,
			BillAmount: v.BillAmount,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
		})
	}
	return res, nil
}

func (c Controller) GetAllTransactionByDate(start string, end string) (*GetAllResponseDataTransaction, error) {
	transaction, err := c.useCase.GetAllTransactionByDate(start, end)
	if err != nil {
		return nil, err
	}

	res := &GetAllResponseDataTransaction{
		Code:      200,
		Message:   "Data Berhasil Diambil",
		Error:     "Success",
		TotalData: len(transaction),
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:         v.Id,
			OdaNumber:  v.OdaNumber,
			BillAmount: v.BillAmount,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
		})
	}
	return res, nil

}

func (c Controller) GetAllTransactionByStatusDate(status string, start string, end string) (*GetAllResponseDataTransaction, error) {
	transaction, err := c.useCase.GetAllTransactionByStatusDate(status, start, end)
	if err != nil {
		return nil, err
	}

	res := &GetAllResponseDataTransaction{
		Code:      http.StatusOK,
		Message:   "Data Berhasil Diambil",
		Error:     "Success",
		TotalData: len(transaction),
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:         v.Id,
			OdaNumber:  v.OdaNumber,
			BillAmount: v.BillAmount,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
		})
	}
	return res, nil
}

func (c Controller) GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) (*GetAllResponseDataTransaction, error) {

	transaction, err := c.useCase.GetTransactionByStatusAndDate(req, input)
	if err != nil {
		return nil, err
	}

	res := &GetAllResponseDataTransaction{
		Code:      200,
		Message:   "Data Berhasil Diambil",
		Error:     "Success",
		TotalData: len(transaction),
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:         v.Id,
			OdaNumber:  v.OdaNumber,
			BillAmount: v.BillAmount,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
		})
	}
	return res, nil
}

func (c Controller) GetTransactionByDate(req FilterByDate, input FilterLimit) (*GetAllResponseDataTransaction, error) {
	transaction, err := c.useCase.GetTransactionByDate(req, input)
	if err != nil {
		return nil, err
	}

	res := &GetAllResponseDataTransaction{
		Code:      200,
		Message:   "Data Berhasil Diambil",
		Error:     "Success",
		TotalData: len(transaction),
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:         v.Id,
			OdaNumber:  v.OdaNumber,
			BillAmount: v.BillAmount,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
		})
	}
	return res, nil

}

func (c Controller) GetTransaction(req *FilterByStatusDate) (*GetAllResponseDataTransaction, error) {

	res := &GetAllResponseDataTransaction{
		Code:    http.StatusBadRequest,
		Message: "Gagal Mengambil Data",
		Error:   "Failed",
		Data:    nil,
	}

	switch {
	case req.Status == "" && req.StartDate == "" && req.EndDate == "":
		return c.GetAllTransaction()
	case req.Status == "":
		return c.GetAllTransactionByDate(req.StartDate, req.EndDate)
	case req.StartDate == "" && req.EndDate == "":
		return c.GetAllTransactionByStatus(req.Status)
	case req.StartDate == "" || req.EndDate == "":
		return res, nil
	default:
		return c.GetAllTransactionByStatusDate(req.Status, req.StartDate, req.EndDate)
	}

}
