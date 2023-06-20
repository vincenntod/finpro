package transactionNotClean

import (
	"golang/auth"
	"golang/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	var transactions []auth.Transaction
	if err := model.DB.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
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
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "4"))
	if err := model.DB.Where("id = ?", id).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}

	model.DB.Model(&transactions).Count(&count)

	if err := model.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
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

func GetAllTransactionByStatusDate(c *gin.Context) {
	var transactions []auth.Transaction
	status := c.Param("status")
	start := c.Param("start")
	end := c.Param("end")

	if err := model.DB.Where("status =? AND(created_at BETWEEN ? AND ?)", status, start, end).Find(&transactions).Error; err != nil {
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

func GetAllTransactionByDate(c *gin.Context) {
	var transactions []auth.Transaction
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

