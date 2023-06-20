package transaction

import (
	"time"
)

type Controller struct {
	useCase UseCaseInterface
}

type ControllerInterface interface {
	GetAllTransaction() (*GetAllResponseDataTransaction, error)
	GetTransactionByStatus(status string) (*GetAllResponseDataTransaction, error)
	GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) (*GetAllResponseDataTransaction, error)
	GetTransactionByDate(req FilterByDate, input FilterLimit) (*GetAllResponseDataTransaction, error)
}

type TransactionItemResponse struct {
	/*
		Id        int       `json:"id"`
		OdaNumber int       `json:"oda_number"`
		Status    string    `json:"status"`
		Price     float32   `json:"price"`
		TotalData int       `json:"total_data"`
		CreatedAt time.Time `json:"created_at"`
	*/

	Id               int       `json:"id"`
	OdaNumber        int       `json:"oda_number"`
	BankAccountNo    int       `json:"bank_account_no"`
	BillingCycleDate string    `json:"billing_cycle_date"`
	PaymentDueDate   time.Time `json:"payment_due_date"`
	OverflowAmount   float32   `json:"overflow_amount"`
	BillAmount       float32   `json:"bill_amount"`
	PrincipalAmount  float32   `json:"principal_amount"`
	InterestAmount   float32   `json:"interest_amount"`
	TotalFeeAmount   float32   `json:"total_fee_amount"`
	Status           string    `json:"status"`
	PaymentMethod    string    `json:"payment_method"`
	AutoDebetCounter int       `json:"auto_debet_counter"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	IsHold           bool      `json:"is_hold"`
	IsFstlPending    bool      `json:"is_fstl_pending"`
	IsHstlPending    bool      `json:"is_hstl_pending"`
	IsLaaPositif     bool      `json:"is_laa_positif"`
	PaymentAmount    float32   `json:"payment_amount"`
	BillingGenDate   time.Time `json:"billing_gen_date"`
	IsOdaPositif     bool      `json:"is_oda_positif"`
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
		Code:    200,
		Message: "Success",
		Error:   "Not Found",
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:               v.Id,
			OdaNumber:        v.OdaNumber,
			BankAccountNo:    v.BankAccountNo,
			BillingCycleDate: v.BillingCycleDate,
			PaymentDueDate:   v.PaymentDueDate,
			OverflowAmount:   v.OverflowAmount,
			BillAmount:       v.BillAmount,
			PrincipalAmount:  v.PrincipalAmount,
			InterestAmount:   v.InterestAmount,
			TotalFeeAmount:   v.TotalFeeAmount,
			Status:           v.Status,
			PaymentMethod:    v.PaymentMethod,
			AutoDebetCounter: v.AutoDebetCounter,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
			IsHold:           v.IsHold,
			IsFstlPending:    v.IsFstlPending,
			IsHstlPending:    v.IsHstlPending,
			IsLaaPositif:     v.IsLaaPositif,
			PaymentAmount:    v.PaymentAmount,
			BillingGenDate:   v.BillingGenDate,
			IsOdaPositif:     v.IsOdaPositif,
		})
	}
	return res, nil
}

func (c Controller) GetTransactionByStatus(status string) (*GetAllResponseDataTransaction, error) {

	transaction, err := c.useCase.GetTransactionByStatus(status)
	if err != nil {
		return nil, err
	}

	res := &GetAllResponseDataTransaction{
		Code:    200,
		Message: "Success",
		Error:   "Not Found",
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:               v.Id,
			OdaNumber:        v.OdaNumber,
			BankAccountNo:    v.BankAccountNo,
			BillingCycleDate: v.BillingCycleDate,
			PaymentDueDate:   v.PaymentDueDate,
			OverflowAmount:   v.OverflowAmount,
			BillAmount:       v.BillAmount,
			PrincipalAmount:  v.PrincipalAmount,
			InterestAmount:   v.InterestAmount,
			TotalFeeAmount:   v.TotalFeeAmount,
			Status:           v.Status,
			PaymentMethod:    v.PaymentMethod,
			AutoDebetCounter: v.AutoDebetCounter,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
			IsHold:           v.IsHold,
			IsFstlPending:    v.IsFstlPending,
			IsHstlPending:    v.IsHstlPending,
			IsLaaPositif:     v.IsLaaPositif,
			PaymentAmount:    v.PaymentAmount,
			BillingGenDate:   v.BillingGenDate,
			IsOdaPositif:     v.IsOdaPositif,
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
		Code:    200,
		Message: "Success",
		Error:   "Not Found",
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:               v.Id,
			OdaNumber:        v.OdaNumber,
			BankAccountNo:    v.BankAccountNo,
			BillingCycleDate: v.BillingCycleDate,
			PaymentDueDate:   v.PaymentDueDate,
			OverflowAmount:   v.OverflowAmount,
			BillAmount:       v.BillAmount,
			PrincipalAmount:  v.PrincipalAmount,
			InterestAmount:   v.InterestAmount,
			TotalFeeAmount:   v.TotalFeeAmount,
			Status:           v.Status,
			PaymentMethod:    v.PaymentMethod,
			AutoDebetCounter: v.AutoDebetCounter,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
			IsHold:           v.IsHold,
			IsFstlPending:    v.IsFstlPending,
			IsHstlPending:    v.IsHstlPending,
			IsLaaPositif:     v.IsLaaPositif,
			PaymentAmount:    v.PaymentAmount,
			BillingGenDate:   v.BillingGenDate,
			IsOdaPositif:     v.IsOdaPositif,
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
		Code:    200,
		Message: "Success",
		Error:   "Not Found",
	}

	for _, v := range transaction {
		res.Data = append(res.Data, TransactionItemResponse{
			Id:               v.Id,
			OdaNumber:        v.OdaNumber,
			BankAccountNo:    v.BankAccountNo,
			BillingCycleDate: v.BillingCycleDate,
			PaymentDueDate:   v.PaymentDueDate,
			OverflowAmount:   v.OverflowAmount,
			BillAmount:       v.BillAmount,
			PrincipalAmount:  v.PrincipalAmount,
			InterestAmount:   v.InterestAmount,
			TotalFeeAmount:   v.TotalFeeAmount,
			Status:           v.Status,
			PaymentMethod:    v.PaymentMethod,
			AutoDebetCounter: v.AutoDebetCounter,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
			IsHold:           v.IsHold,
			IsFstlPending:    v.IsFstlPending,
			IsHstlPending:    v.IsHstlPending,
			IsLaaPositif:     v.IsLaaPositif,
			PaymentAmount:    v.PaymentAmount,
			BillingGenDate:   v.BillingGenDate,
			IsOdaPositif:     v.IsOdaPositif,
		})
	}
	return res, nil

}
