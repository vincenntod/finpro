package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

type RequestHandlerinterface interface {
	GetAllTransaction(c *gin.Context)
	GetTransactionByStatus(c *gin.Context)
	GetTransactionByStatusAndDate(c *gin.Context)
	GetTransactionByDate(c *gin.Context)
	GetAllTransactionByDate(c *gin.Context)
}

type GetAllResponseDataTransaction struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Error   string                    `json:"err"`
	Data    []TransactionItemResponse `json:"data"`
}

type FilterByDate struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type FilterByStatusDate struct {
	Status    string `json:"status"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type FilterLimit struct {
	Page     int
	PageSize int
}

func NewRequestHandler(ctrl ControllerInterface) RequestHandlerinterface {
	return RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) RequestHandlerinterface {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

func (h RequestHandler) GetAllTransaction(c *gin.Context) {

	res, err := h.ctrl.GetAllTransaction()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) GetTransactionByStatus(c *gin.Context) {
	status := c.Param("status")

	if status == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "status not found"})
		return
	}

	res, err := h.ctrl.GetTransactionByStatus(status)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) GetAllTransactionByDate(c *gin.Context) {
	start := c.Param("start")
	end := c.Param("end")

	if start == "" || end == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "status not found"})
		return
	}

	res, err := h.ctrl.GetAllTransactionByDate(start, end)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) GetTransactionByStatusAndDate(c *gin.Context) {
	var req FilterByStatusDate

	parameterPage := c.Param("page")
	page, _ := strconv.Atoi(c.DefaultQuery("page", parameterPage))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))

	input := FilterLimit{
		Page:     page,
		PageSize: pageSize,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := h.ctrl.GetTransactionByStatusAndDate(req, input)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) GetTransactionByDate(c *gin.Context) {
	var req FilterByDate

	parameterPage := c.Param("page")
	page, _ := strconv.Atoi(c.DefaultQuery("page", parameterPage))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))

	input := FilterLimit{
		Page:     page,
		PageSize: pageSize,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := h.ctrl.GetTransactionByDate(req, input)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}
