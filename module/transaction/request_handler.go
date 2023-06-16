package transaction

import (
	
	"net/http"
	

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	ctrl *Controller
}

func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}



// func (h RequestHandler) GetAllTransaction(c *gin.Context) {
	
// 	id := c.Param("id")

// 	if id == "" {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
// 		return
// 	}
	
// 	res, err := h.ctrl.GetAllTransaction()
// 	if err != nil {
// 		c.JSON(500, gin.H{"message": "Data Tidak Ditemukan"})
// 		return
// 	}
// 	c.JSON(200, res)

// }

func (h RequestHandler) GetTransactionByStatus(c *gin.Context) {
	status := c.Param("status")

	if status == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
		return
	}
	
	res, err := h.ctrl.GetTransactionByStatus(status)
	if err != nil {
		c.JSON(500, gin.H{"message": "Data Tidak Ditemukan"})
		return
	}
	c.JSON(200, res)

}