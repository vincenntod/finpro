package transaction

import (
	"time"
)

type Controller struct {
	useCase *UseCase
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

type GetAllResponseByStatus struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Error   string                  `json:"error"`
	Data    TransactionItemResponse `json:"data"`
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

// func (c Controller) GetAllTransaction() (*GetAllResponseByStatus, error) {
// }

func (c Controller) GetTransactionByStatus(status string) (*GetAllResponseByStatus, error) {

	transaction, err := c.useCase.GetTransactionByStatus(status)
	if err != nil {
		return nil, err
	}

	res := &GetAllResponseByStatus{
		Code:    200,
		Message: "Success",
		Error:   "Not Found",
		Data: TransactionItemResponse{
			Id:               transaction.Id,
			OdaNumber:        transaction.OdaNumber,
			BankAccountNo:    transaction.BankAccountNo,
			BillingCycleDate: transaction.BillingCycleDate,
			PaymentDueDate:   transaction.PaymentDueDate,
			OverflowAmount:   transaction.OverflowAmount,
			BillAmount:       transaction.BillAmount,
			PrincipalAmount:  transaction.PrincipalAmount,
			InterestAmount:   transaction.InterestAmount,
			TotalFeeAmount:   transaction.TotalFeeAmount,
			Status:           transaction.Status,
			PaymentMethod:    transaction.PaymentMethod,
			AutoDebetCounter: transaction.AutoDebetCounter,
			CreatedAt:        transaction.CreatedAt,
			UpdatedAt:        transaction.UpdatedAt,
			IsHold:           transaction.IsHold,
			IsFstlPending:    transaction.IsFstlPending,
			IsHstlPending:    transaction.IsHstlPending,
			IsLaaPositif:     transaction.IsLaaPositif,
			PaymentAmount:    transaction.PaymentAmount,
			BillingGenDate:   transaction.BillingGenDate,
			IsOdaPositif:     transaction.IsOdaPositif,
		},
	}
	return res, nil
}
