package account

import (
	"golang/module/account/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandlerInterface interface {
	GetDataUser(c *gin.Context)
	GetDataUserById(c *gin.Context)
	CreateAccount(c *gin.Context)
	EditDataUser(c *gin.Context)
	DeleteDataUser(c *gin.Context)
	Login(c *gin.Context)
	SendEmailForgotPassword(c *gin.Context)
	SendEmailRegistration(c *gin.Context)
	CompareVerificationCode(c *gin.Context)
	EditPassword(c *gin.Context)
}
type RequestHandler struct {
	Ctrl ControllerInterface
}

func NewRequestHandler(ctrl ControllerInterface) RequestHandlerInterface {
	return &RequestHandler{
		Ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) RequestHandlerInterface {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

func (h RequestHandler) GetDataUser(c *gin.Context) {
	res, err := h.Ctrl.GetDataUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) GetDataUserById(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Ctrl.GetDataUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) CreateAccount(c *gin.Context) {
	var account dto.CreateRequest
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	res, _ := h.Ctrl.CreateAccount(&account)
	c.JSON(http.StatusOK, res)

}

func (h RequestHandler) EditDataUser(c *gin.Context) {
	var account dto.EditDataUserRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.Ctrl.EditDataUser(id, &account)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) DeleteDataUser(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Ctrl.DeleteDataUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) Login(c *gin.Context) {
	var account dto.LoginResponseRequest

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	token, res, err := h.Ctrl.Login(&account)
	if err != nil {
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	c.Writer.Header().Set("Authorization", token)
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) SendEmailForgotPassword(c *gin.Context) {
	email := c.Param("email")
	res, err := h.Ctrl.SendEmailForgotPassword(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) SendEmailRegistration(c *gin.Context) {
	email := c.Param("email")
	res, err := h.Ctrl.SendEmailRegistration(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) CompareVerificationCode(c *gin.Context) {
	var verificationCodeRequest dto.VerificationCodeRequest

	if err := c.ShouldBindJSON(&verificationCodeRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Internal Server Error",
		})
		return
	}
	res, err := h.Ctrl.CompareVerificationCode(&verificationCodeRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) EditPassword(c *gin.Context) {
	var req dto.EditDataUserRequest
	queryUrl := c.Request.URL.Query()
	id := queryUrl.Get("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Bad Request Id Parameter kosong",
		})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Internal Server Error",
		})
		return
	}

	verificationCode := c.Param("verification-code")
	code := verificationCode
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Bad Request Verification Code nil",
		})
		return
	}

	res, err := h.Ctrl.EditPassword(id, code, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
	}
	c.JSON(http.StatusOK, res)
}
