package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	Ctrl ControllerInterface
}

type RequestHandlerinterface interface {
	GetAllTransaction(c *gin.Context)
	GetAllTransactionByStatus(c *gin.Context)
	GetAllTransactionByDate(c *gin.Context)
	GetAllTransactionByStatusDate(c *gin.Context)
	GetAllLimit(c *gin.Context)
}

type MessageResponse struct {
	Message string `json:"message"`
}

type GetAllResponseDataTransaction struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	Error     string                    `json:"status"`
	TotalData int                       `json:"total_data"`
	Data      []TransactionItemResponse `json:"data"`
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

func FormatDate(date string) string {
	day := string(date[:2])
	month := string(date[3:5])
	year := string(date[6:10])

	parseDate := year + "-" + month + "-" + day
	return parseDate
}

func NewRequestHandler(ctrl ControllerInterface) RequestHandlerinterface {
	return RequestHandler{
		Ctrl: ctrl,
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

	res, err := h.Ctrl.GetAllTransaction()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) GetAllTransactionByStatus(c *gin.Context) {
	status := c.Param("status")

	if status == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "status not found"})
		return
	}

	res, err := h.Ctrl.GetAllTransactionByStatus(status)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) GetAllTransactionByDate(c *gin.Context) {
	start := FormatDate(c.Param("start"))
	end := FormatDate(c.Param("end"))

	res, err := h.Ctrl.GetAllTransactionByDate(start, end)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) GetAllTransactionByStatusDate(c *gin.Context) {
	status := c.Param("status")
	start := FormatDate(c.Param("start"))
	end := FormatDate(c.Param("end"))

	if status == "" || start == "" || end == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "inputan tidak boleh kosong"})
		return
	}

	res, err := h.Ctrl.GetAllTransactionByStatusDate(status, start, end)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)

}

func (h RequestHandler) GetAllLimit(c *gin.Context) {
	parameterPage := c.Param("id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", parameterPage))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))

	input := FilterLimit{
		Page:     page,
		PageSize: pageSize,
	}

	if parameterPage == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "parameter not found"})
		return
	}

	res, err, _ := h.Ctrl.GetAllLimit(input)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, res)
}
