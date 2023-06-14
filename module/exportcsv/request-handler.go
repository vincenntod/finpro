package exportcsv

import (
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExpoertCSVRequestHandler struct {
	ctrl *ExportCSVController
}

func NewRequestHandler(ctrl *ExportCSVController) *ExpoertCSVRequestHandler {
	return &ExpoertCSVRequestHandler{
		ctrl: ctrl,
	}

}

func DefaultRequestHandler(db *gorm.DB) *ExpoertCSVRequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type ExportCSVRequest struct {
	Status string `json:"status"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func (h ExpoertCSVRequestHandler) ExportCSVHandler(c *gin.Context) {

	var req ExportCSVRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	exportData, err := h.ctrl.ExportCSV(&req)
	switch {
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
