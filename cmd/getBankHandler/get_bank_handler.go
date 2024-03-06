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

// GetBankByIDHandler handler for GetBankByIDHandler new bank.
func (h *BankHandler) GetBankByIDHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("GetBankByIDHandler", ctx, request)
	createBankHandler, errorBankHandler := h.bankService.GetBankByIDService(ctx, request)
	if errorBankHandler != nil {
		logs.LogTrackingError("GetBankByIDHandler", "GetBankByIDService", ctx, request, errorBankHandler)
		return response.ErrorResponse(http.StatusInternalServerError, constantsmicro.ErrorGettingBank)
	}
	return createBankHandler, nil
}

func main() {
	// Create an instance of BankHandler in the main function.
	bankHandler := NewBankHandler()

	// Wrap the handler function with logging middleware.
	handlerWithLogging := metadata.MiddlewareMetadata(bankHandler.GetBankByIDHandler)

	// Start the Lambda handler with the handler function wrapped in the middleware.
	lambda.Start(handlerWithLogging)
}
