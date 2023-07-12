package exportcsv

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"

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

type ExportCSVRequest struct {
	Status    string `form:"status"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
type UrlRequest struct {
	ClientUrl string
	PathUrl   string
}

type ErrorResponse struct {
	Error string `json:"error"`
}
type MessageResponse struct {
	Message string `json:"message"`
}

func (h RequestHandler) ExportCSVHandler(c *gin.Context) {

	var req ExportCSVRequest
	var url UrlRequest
	err := c.Bind(&req)
	fileType := c.Param("type")
	url.ClientUrl = c.Request.URL.String()
	url.PathUrl = c.Request.URL.Path

	fmt.Println(url.ClientUrl, " ", url.PathUrl)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})

		return
	}
	serverGot := url.PathUrl + "?" + "start_date=" + req.StartDate + "&end_date=" + req.EndDate
	if serverGot == url.ClientUrl {
		fmt.Println(serverGot, " ][", url.ClientUrl)
	}
	exportData, err := h.Ctrl.ExportCSV(&req, &url)
	switch {
	case err != nil && err.Error() == "Not Found":
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
	case err != nil && err.Error() == "Invalid Request":
		c.JSON(http.StatusNotAcceptable, ErrorResponse{Error: err.Error()})
	case err != nil:
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	switch fileType {
	case "csv":
		GenerateCSV(c, exportData)
	case "xlsx":
		GenerateXLSX(c, exportData)
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "export transaction success"})

}
func GenerateCSV(c *gin.Context, data [][]string) {
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=transaction-export.csv")
	writer := csv.NewWriter(c.Writer)
	if err := writer.WriteAll(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate CSV file"})
	}
}

func GenerateXLSX(c *gin.Context, data [][]string) {
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=transaction-export.xlsx")
	// Create a new Excel file
	file := excelize.NewFile()

	// Write the data to the Excel file
	for rowIndex, row := range data {
		for colIndex, cellValue := range row {
			cell := excelize.ToAlphaString(colIndex+1) + fmt.Sprint(rowIndex+1)
			file.SetCellValue("sheet1", cell, cellValue)

		}
	}
	file.Write(c.Writer)
}
