package test

import (
	"golang/module/transaction"
	mocks "golang/module/transaction/mocks/module/transaction"

	"testing"

	"github.com/gin-gonic/gin"
)

func TestRequestHandler_GetAllLimit(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Limit",
			fields: fields{
				ctrl: mocks.NewControllerInterface(t),
			},
			args: args{
				c: &gin.Context{
					Params: gin.Params{
						gin.Param{
							Key:   "id",
							Value: "1",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := transaction.RequestHandler{
				Ctrl: tt.fields.ctrl,
			}
		/*	tt.fields.ctrl.(*mocks.ControllerInterface).EXPECT().GetAllLimit(tt.args.c).Return([]transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-01 00:00:00",
				},
			}, nil, 0)*/
			h.GetAllLimit(tt.args.c)
		})
	}
}

func TestRequestHandler_GetAllTransaction(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := transaction.RequestHandler{
				Ctrl: tt.fields.ctrl,
			}
			
			h.GetAllTransaction(tt.args.c)
		})
	}
}

func TestRequestHandler_GetAllTransactionByDate(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := transaction.RequestHandler{
				Ctrl: tt.fields.ctrl,
			}
			h.GetAllTransactionByDate(tt.args.c)
		})
	}
}

func TestRequestHandler_GetAllTransactionByStatus(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := transaction.RequestHandler{
				Ctrl: tt.fields.ctrl,
			}
			h.GetAllTransactionByStatus(tt.args.c)
		})
	}
}

func TestRequestHandler_GetAllTransactionByStatusDate(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := transaction.RequestHandler{
				Ctrl: tt.fields.ctrl,
			}
			h.GetAllTransactionByStatusDate(tt.args.c)
		})
	}
}
