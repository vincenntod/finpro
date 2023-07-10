package helper

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateTestServer(request *http.Request, route func(router *gin.Engine)) (int, []byte) {
	gin.SetMode(gin.TestMode)
	responseRecorder := httptest.NewRecorder()
	_, router := gin.CreateTestContext(responseRecorder)

	route(router)
	router.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := io.ReadAll(response.Body)
	defer func() {
		_ = response.Body.Close()
	}()

	return response.StatusCode, body

}

func NewMockQueryDb(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	conn, mockQuery, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mysqlConfig := mysql.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	}
	options := gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}
	mockDb, _ := gorm.Open(mysql.New(mysqlConfig), &options)

	return mockQuery, mockDb
}
