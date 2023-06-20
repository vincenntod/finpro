package transactions

import (
	"golang/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
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

func GetAllTransactions(c *gin.Context) {
	var transactions []Transaction
	if err := model.DB.Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error"})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Success",
		"error":   "Tidak Ada Error",
		"data":    &transactions,
	})
}

func GetAllTransactionsLimit(c *gin.Context) {
	var transactions []Transaction
	var count int64
	id := c.Param("id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", id))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "100"))
	if err := model.DB.Where("id = ?", id).Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error"})
		return
	}

	model.DB.Model(&transactions).Count(&count)

	if err := model.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error"})
		return
	}
	c.JSON(200, gin.H{
		"code":              200,
		"message":           "Success",
		"error":             "Tidak Ada Error",
		"data":              &transactions,
		"Jumlah Semua Data": count,
	})
}

func GetAllTransactionByStatus(c *gin.Context) {
	var transactions []Transaction
	status := c.Param("status")
	if err := model.DB.Where("status = ?", status).Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error"})
		return

	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Success",
		"error":   "Tidak Ada Error",
		"data":    &transactions,
	})
}

/*type FilterByStatus struct {
	Status string `json:"status"`
}

func GetTransactionByStatus(c *gin.Context) {
	var transactions []Transaction
	status := c.Param("status")
	if err := model.DB.Where("status = ?", status).Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error"})
		return

	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Success",
		"error":   "Tidak Ada Error",
		"data":    &transactions,
	})
}
*/

type FilterByDate struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func GetAllTransactionByDate(c *gin.Context) {
	var transactions []Transaction
	start := c.Param("start")
	end := c.Param("end")

	if err := model.DB.Where("created_at BETWEEN ? AND ?", start, end).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Success",
		"error":   "Tidak Ada Error",
		"data":    &transactions,
	})
}

type FilterByStatusDate struct {
	Status    string `json:"status"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func GetTransactionByStatusDate(c *gin.Context) {
	var transactions []Transaction
	var filterByStatusDate FilterByStatusDate

	parameterPage := c.Param("page")
	page, _ := strconv.Atoi(c.DefaultQuery("page", parameterPage))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))

	if err := c.ShouldBindJSON(&filterByStatusDate); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	if err := model.DB.Offset((page-1)*pageSize).Limit(pageSize).Order("created_at desc").Where("status =? AND(created_at BETWEEN ? AND ?)", filterByStatusDate.Status, filterByStatusDate.StartDate, filterByStatusDate.EndDate).Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"status":  "Success",
		"page":    &parameterPage,
		"message": "Success Get Transaction",
		"data":    &transactions})

}
