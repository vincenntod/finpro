package account

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

type CreateRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type LoginResponseRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
type EditDataUserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h RequestHandler) GetDataUser(c *gin.Context) {
	res, err := h.ctrl.GetDataUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, res)
}
func (h RequestHandler) GetDataUserById(c *gin.Context) {
	id := c.Param("id")
	res, err := h.ctrl.GetDataUserById(id)
	if err != nil {
		c.JSON(500, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, res)
}
func (h RequestHandler) CreateAccount(c *gin.Context) {
	var account CreateRequest
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	res, _ := h.ctrl.CreateAccount(&account)
	c.JSON(200, res)

}

func (h RequestHandler) EditDataUser(c *gin.Context) {
	var account EditDataUserRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.ctrl.EditDataUser(id, &account)
	if err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, res)
}
func (h RequestHandler) DeleteDataUser(c *gin.Context) {
	id := c.Param("id")
	res, err := h.ctrl.DeleteDataUser(id)
	if err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
	}
	c.JSON(200, res)
}
func (h RequestHandler) Login(c *gin.Context) {
	var account LoginResponseRequest

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	token, res, err := h.ctrl.Login(&account)
	if err != nil {
		c.JSON(401, res)
		return
	}
	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, res)
}
func (h RequestHandler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Berhasil Logout"})
}
func (h RequestHandler) SendEmail(c *gin.Context) {
	email := c.Param("email")
	res, err := h.ctrl.SendEmail(email)
	if err != nil {
		c.JSON(401, res)
	}
	c.JSON(200, res)

}
func (h RequestHandler) CompareVerificationCode(c *gin.Context) {
	var verificationCodeRequest VerificationCodeRequest

	if err := c.ShouldBindJSON(&verificationCodeRequest); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Internal Server Error",
		})
		return
	}
	res, err := h.ctrl.CompareVerificationCode(&verificationCodeRequest)
	if err != nil {
		c.JSON(400, res)
		return
	}
	c.JSON(200, res)
}
func (h RequestHandler) EditPassword(c *gin.Context) {
	var req EditDataUserRequest
	queryUrl := c.Request.URL.Query()
	id := queryUrl.Get("id")
	if id == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Bad Request id nil",
		})
		return
	}

	code := c.Request.Header.Get("VerificationCode")
	if code == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Bad Request Verification Code nil",
		})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Internal Server Error",
		})
		return
	}
	res, err := h.ctrl.EditPassword(id, code, &req)
	if err != nil {
		c.JSON(500, res)
	}
	c.JSON(200, res)
}
