package transaction

import (
	"golang/module/transaction/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	Ctrl ControllerInterface
}

type RequestHandlerinterface interface {
	GetAllTransactionByRequestLimit(c *gin.Context)
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

func (h RequestHandler) GetAllTransactionByRequestLimit(c *gin.Context) {
	var input dto.Request
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := h.Ctrl.GetAllTransactionByRequest(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
