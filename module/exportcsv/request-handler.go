package exportcsv

import (
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExportCSVRequestHandler struct {
	ctrl ExportCSVController
}

func NewRequestHandler(ctrl ExportCSVController) *ExportCSVRequestHandler {
	return &ExportCSVRequestHandler{
		ctrl: ctrl,
	}

}

func DefaultRequestHandler(db *gorm.DB) *ExportCSVRequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type ExportCSVRequest struct {
	Status    string `form:"status"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func (h ExportCSVRequestHandler) ExportCSVHandler(c *gin.Context) {

	var req ExportCSVRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	exportData, err := h.ctrl.ExportCSV(&req)
	switch {
	case err != nil && err.Error() == "Not Found":
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
	case err != nil && err.Error() == "invalid field status":
		c.JSON(http.StatusNotAcceptable, ErrorResponse{Error: err.Error()})
	case err != nil:
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=transaction-export.csv")
	writer := csv.NewWriter(c.Writer)
	if err := writer.WriteAll(exportData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate CSV file"})
	}

}
