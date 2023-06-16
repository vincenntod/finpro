package transaction

import (
	"golang/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id        int       `json:"id"`
	OdaNumber int       `json:"oda_number"`
	Status    string    `json:"status"`
	Price     float32   `json:"price"`
	TotalData int       `json:"total_data"`
	CreatedAt time.Time `json:"created_at"`
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
		"data":    &transactions})
}

func GetAllTransactionsByParam(c *gin.Context) {
	var transactions []Transaction
	var count int64
	id := c.Param("id")
	end := c.Param("end")
	page, _ := strconv.Atoi(c.DefaultQuery("page", id))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", end))
	if err := model.DB.Where("id = ?", id).Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error"})
		return
	}

	model.DB.Model(&transactions).Count(&count)
	model.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&transactions)
	if err := model.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"message": "Error"})
		return
	}
	c.JSON(200, gin.H{
		"code":            200,
		"message":         "Success",
		"error":           "Tidak Ada Error",
		"data":            &transactions,
		"Jumlah All Page": count,
	})
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

func FindByDate(c *gin.Context) {
	var transactions []Transaction
	date := c.Param("date")
	if err := model.DB.Where("created_at = ?", date).Find(&transactions).Error; err != nil {
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
