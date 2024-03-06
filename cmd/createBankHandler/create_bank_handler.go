package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/constantsmicro"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/service"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"github.com/diegocabrera89/ms-payment-core/middleware/metadata"
	"github.com/diegocabrera89/ms-payment-core/response"
	"net/http"
)

// BankHandler handles HTTP requests related to the bank entity.
type BankHandler struct {
	bankService *service.BankService
}

// NewBankHandler create a new Bank Handler instance.
func NewBankHandler() *BankHandler {
	return &BankHandler{
		bankService: service.NewBankService(),
	}
}

// CreateBankHandler handler for CreateBankHandler new bank.
func (h *BankHandler) CreateBankHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("CreateBankHandler", ctx, request)
	createBankHandler, errorBankHandler := h.bankService.CreateBankService(ctx, request)
	if errorBankHandler != nil {
		logs.LogTrackingError("CreateBankHandler", "CreateBankService", ctx, request, errorBankHandler)
		return response.ErrorResponse(http.StatusInternalServerError, constantsmicro.ErrorCreatingBank)
	}
	return createBankHandler, nil
}

func main() {
	// Create an instance of BankHandler in the main function.
	bankHandler := NewBankHandler()

	// Wrap the handler function with logging middleware.
	handlerWithLogging := metadata.MiddlewareMetadata(bankHandler.CreateBankHandler)

	// Start the Lambda handler with the handler function wrapped in the middleware.
	lambda.Start(handlerWithLogging)
}
