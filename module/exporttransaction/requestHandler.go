package exporttransaction

import (
	"encoding/csv"
	"golang/module/exporttransaction/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	Ctrl Controller
}

func NewRequestHandler(ctrl Controller) *RequestHandler {
	return &RequestHandler{
		Ctrl: ctrl,
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

func (h RequestHandler) ExportTransactionHandler(c *gin.Context) {

	var filter dto.Filter
	var url dto.Url
	var req dto.ExportCSVRequest

	err := c.Bind(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	url.ClientUrl = c.Request.URL.String()
	url.PathUrl = c.Request.URL.Path
	req = dto.ExportCSVRequest{Filter: filter, Url: url}

	exportData, err := h.Ctrl.ExportCSV(&req)
	switch {
	case err != nil && err.Error() == "not found":
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: err.Error()})
	case err != nil && err.Error() == "invalid request":
		c.JSON(http.StatusNotAcceptable, dto.ErrorResponse{Error: err.Error()})
	case err != nil:
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=transaction-export.csv")
	writer := csv.NewWriter(c.Writer)
	if err := writer.WriteAll(exportData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate CSV file"})
	}
}
