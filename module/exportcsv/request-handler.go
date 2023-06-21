package exportcsv

import (
	"encoding/csv"
	"golang/module/transactions"
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
	Status    string
	StartDate string
	EndDate   string
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func (h ExpoertCSVRequestHandler) ExportCSVHandler(c *gin.Context) {
<<<<<<< HEAD
	exportData, err := h.ctrl.ExportCSV()
	if err != nil {
		switch {
		case err.Error() == "Not Found":
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		case err.Error() == "Invalid field status":
			c.JSON(http.StatusNotAcceptable, ErrorResponse{Error: err.Error()})
		case err.Error() == "Failed to generate CSV file":
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}
	}
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=transaction-export.csv")
	writer := csv.NewWriter(c.Writer)
	if err := writer.WriteAll(exportData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate CSV file"})
	}

}

func (h ExpoertCSVRequestHandler) ExportCSVHandlerStatusfilter(c *gin.Context) {
	status := c.Param("status")
	exportData, err := h.ctrl.ExportCSVStatusFilter(status)
	if err != nil {
		switch {
		case err.Error() == "Not Found":
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		case err.Error() == "Invalid field status":
			c.JSON(http.StatusNotAcceptable, ErrorResponse{Error: err.Error()})
		case err.Error() == "Failed to generate CSV file":
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}
	}
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=transaction-export.csv")
	writer := csv.NewWriter(c.Writer)
	if err := writer.WriteAll(exportData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate CSV file"})
	}
}

func (h ExpoertCSVRequestHandler) ExportCSVHandlerRangeDateFilter(c *gin.Context) {

	startDate := transactions.FormatDate(c.Param("start_date"))
	endDate := transactions.FormatDate(c.Param("end_date"))
	exportData, err := h.ctrl.ExportCSVRangeDateFilter(startDate, endDate)
	if err != nil {
		switch {
		case err.Error() == "Not Found":
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		case err.Error() == "Invalid field status":
			c.JSON(http.StatusNotAcceptable, ErrorResponse{Error: err.Error()})
		case err.Error() == "Failed to generate CSV file":
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}
	}
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=transaction-export.csv")
	writer := csv.NewWriter(c.Writer)
	if err := writer.WriteAll(exportData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate CSV file"})
	}

}

func (h ExpoertCSVRequestHandler) ExportCSVHandlerStatusAndRangeDateFilter(c *gin.Context) {
	status := c.Param("status")
	startDate := transactions.FormatDate(c.Param("start_date"))
	endDate := transactions.FormatDate(c.Param("end_date"))
	exportData, err := h.ctrl.ExportCSVStatusAndRangeDate(status, startDate, endDate)
	if err != nil {
		switch {
		case err.Error() == "Not Found":
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		case err.Error() == "Invalid field status":
			c.JSON(http.StatusNotAcceptable, ErrorResponse{Error: err.Error()})
		case err.Error() == "Failed to generate CSV file":
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}
=======

	var req ExportCSVRequest
	if err := c.BindJSON(&req); err != nil {
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
>>>>>>> parent of 007e65b (solve error and adding custom resspone invalid json format)
	}

	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=transaction-export.csv")
	writer := csv.NewWriter(c.Writer)
	if err := writer.WriteAll(exportData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate CSV file"})
	}

}
