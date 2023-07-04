package test

import (
	mocks "golang/mocks/module/exportcsv"
	"golang/module/exportcsv"
	"golang/module/exportcsv/helper"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExportCSVHandler(t *testing.T) {
	mockController := mocks.NewExportCSVController(t)
	mockController.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"", "", ""}).Return(
		[][]string{{"id", "od_number", "status", "price", "created_at"},
			{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
			{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
			{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
			{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		nil).Once()

	handler := exportcsv.NewRequestHandler(mockController)
	makeRequest := func() *http.Request {
		request, _ := http.NewRequest(http.MethodGet, "/export-transaction", nil)
		return request
	}

	statusCode, body := helper.CreateTestServer(makeRequest(), func(router *gin.Engine) {
		router.GET("/export-transaction", handler.ExportCSVHandler)
	})
	assert.Equal(t, 200, statusCode)
	assert.NotEmpty(t, body)

}
