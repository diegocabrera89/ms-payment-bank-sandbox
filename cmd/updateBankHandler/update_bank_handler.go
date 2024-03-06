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

// BankHandler handles HTTP requests related to the Bank entity.
type BankHandler struct {
	bankService *service.BankService
}

// NewBankHandler create a new BankHandler instance.
func NewBankHandler() *BankHandler {
	return &BankHandler{
		bankService: service.NewBankService(),
	}
}

// UpdateBankHandler handler for update pet.
func (h *BankHandler) UpdateBankHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("UpdateBankHandler", ctx, request)
	deletePetHandler, err := h.bankService.UpdateAmountBankService(ctx, request)
	if err != nil {
		logs.LogTrackingError("UpdateBankHandler", "UpdateAmountBankService", ctx, request, err)
		return response.ErrorResponse(http.StatusInternalServerError, constantsmicro.ErrorUpdatingBank)
	}
	return deletePetHandler, nil
}

func main() {
	// Create an instance of PetHandler in the main function.
	bankHandler := NewBankHandler()

	// Wrap the handler function with logging middleware.
	handlerWithLogging := metadata.MiddlewareMetadata(bankHandler.UpdateBankHandler)

	// Start the Lambda handler with the handler function wrapped in the middleware.
	lambda.Start(handlerWithLogging)
}
