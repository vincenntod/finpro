package transactionNotClean

import (
	"golang/auth"
	"golang/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	var transactions []auth.Transaction
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
	var transactions []auth.Transaction
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
	var transactions []auth.Transaction
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
	var transactions []auth.Transaction
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

func GetTransactionByDate(c *gin.Context) {
	var transactions []auth.Transaction
	var filterByDate FilterByDate

	parameterPage := c.Param("page")
	page, _ := strconv.Atoi(c.DefaultQuery("page", parameterPage))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))

	if err := c.ShouldBindJSON(&filterByDate); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	if err := model.DB.Offset((page-1)*pageSize).Limit(pageSize).Order("created_at desc").Where("created_at BETWEEN ? AND ?", filterByDate.StartDate, filterByDate.EndDate).Find(&transactions).Error; err != nil {
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

type FilterByStatusDate struct {
	Status    string `json:"status"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func GetTransactionByStatusDate(c *gin.Context) {
	var transactions []auth.Transaction
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