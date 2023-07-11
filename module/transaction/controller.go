package transaction

import (
	"net/http"
)

type Controller struct {
	UseCase UseCaseInterface
}

type ControllerInterface interface {
	GetAllTransaction() (*GetAllResponseDataTransaction, error)
	GetAllTransactionByStatus(status string) (*GetAllResponseDataTransaction, error)
	GetAllTransactionByDate(start string, end string) (*GetAllResponseDataTransaction, error)
	GetAllTransactionByStatusDate(status string, start string, end string) (*GetAllResponseDataTransaction, error)
	GetAllLimit(input FilterLimit) (*GetAllResponseDataTransaction, error, int64)
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
		UseCase: useCase,
	}
}

func (c Controller) GetAllTransaction() (*GetAllResponseDataTransaction, error) {

	transaction, err := c.UseCase.GetAllTransaction()
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

	transaction, err := c.UseCase.GetAllTransactionByStatus(status)
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
	transaction, err := c.UseCase.GetAllTransactionByDate(start, end)
	if err != nil {
		res := &GetAllResponseDataTransaction{
			Code:      400,
			Message:   "Bad Request",
			Error:     "Error",
			TotalData: len(transaction),
		}
		return res, nil
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
	transaction, err := c.UseCase.GetAllTransactionByStatusDate(status, start, end)
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

func (c Controller) GetAllLimit(input FilterLimit) (*GetAllResponseDataTransaction, error, int64) {
	transaction, err, total := c.UseCase.GetAllLimit(input)
	if err != nil {
		return nil, err, 0
	}

	res := &GetAllResponseDataTransaction{
		Code:      200,
		Message:   "Data Berhasil Diambil",
		Error:     "Success",
		TotalData: int(total),
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
	return res, nil, total
}
