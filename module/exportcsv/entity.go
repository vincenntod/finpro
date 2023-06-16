package exportcsv

import (
	"time"
)

type Transaction struct {
	Id                 int       `json: "id"`
	Oda_number         string    `json: "od_number" `
	Bank_account_no    string    `json: "bank_acount_no"`
	Billing_cycle_date string    `json: "billing_cycle_date"`
	Payment_due_date   time.Time `json: "payment_due_date" `
	Overflow_amount    float64   `json: "overflow_amount"`
	Bill_amount        float64   `json: "bill_amount"`
	Principal_amount   float64   `json: "principal_amount"`
	Interest_amount    float64   `json: "interest_amount"`
	Total_fee_amount   float64   `json: "total_fee_amount"`
	Status             string    `json: "status"`
	Payment_method     string    `json: "payment_method"`
	Auto_debet_counter int       `json: "auto_debet_counter"`
	Created_at         time.Time `json: "created_at"`
	Updated_at         time.Time `json: "updated_at"`
	Is_hold            bool      `json: "is_hold"`
	Is_fstl_pending    bool      `json: "is_fstl_pending"`
	Is_hstl_pending    bool      `json: "is_hstl_pending"`
	Is_laa_positif     bool      `json: "is_laa_positif"`
	Payment_amount     float64   `json: "payment_amount"`
	Billing_gen_date   time.Time `json: "billing_gen_date"`
	Is_oda_positif     bool      `json: "is_oda_positif"`
}
