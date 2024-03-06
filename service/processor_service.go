package service

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/constantsmicro"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/models"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/repository"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/utils"
	"github.com/diegocabrera89/ms-payment-core/constantscore"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"github.com/diegocabrera89/ms-payment-core/response"
	"net/http"
	"time"
)

// BankService represents the service for the BankService entity.
type BankService struct {
	bankRepo *repository.BankRepositoryImpl
}

// NewBankService create a new BankService instance.
func NewBankService() *BankService {
	return &BankService{
		bankRepo: repository.NewBankRepository(),
	}
}

// CreateBankService handles the creation of a new bank.
func (r *BankService) CreateBankService(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("CreateBankService", ctx, request)
	var bank models.Bank
	err := json.Unmarshal([]byte(request.Body), &bank)
	if err != nil {
		logs.LogTrackingError("CreateBankService", "JSON Unmarshal", ctx, request, err)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.InvalidRequestBody)
	}

	utils.BuildBank(&bank)

	createBank, errorBankRepository := r.bankRepo.CreateBankRepository(ctx, request, bank)
	if errorBankRepository != nil {
		logs.LogTrackingError("CreateBankService", "CreateBankRepository", ctx, request, errorBankRepository)
		return response.ErrorResponse(http.StatusInternalServerError, constantscore.ErrorCreatingItem)
	}

	responseBody, err := json.Marshal(createBank)
	if err != nil {
		logs.LogTrackingError("CreateBankService", "JSON Marshal", ctx, request, err)
		return response.ErrorResponse(http.StatusInternalServerError, constantscore.InvalidResponseBody)
	}
	return response.SuccessResponse(http.StatusCreated, responseBody, constantscore.ItemCreatedSuccessfully)
}

// GetBankByIDService handles the creation of a new bank.
func (r *BankService) GetBankByIDService(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("GetBankByIDService", ctx, request)
	logs.LogTrackingInfoData("GetBankByIDService request", request, ctx, request)
	var responseBody []byte
	bankID := request.PathParameters[constantsmicro.BankID]
	logs.LogTrackingInfoData("GetBankByIDService bankID", bankID, ctx, request)
	if bankID == "" {
		logs.LogTrackingError("GetBankByIDService", "PathParameters", ctx, request, nil)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.ErrorGettingElement)
	}

	getBankByID, err := r.bankRepo.GetBankByBankIDRepository(ctx, request, bankID)
	if err != nil {
		logs.LogTrackingError("GetBankByIDService", "GetBankByBankIDRepository", ctx, request, err)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.ErrorGettingElement)
	}

	if getBankByID.BankID != "" {
		responseBody, err = json.Marshal(getBankByID)
		if err != nil {
			logs.LogTrackingError("GetBankByIDService", "JSON Marshal", ctx, request, err)
			return response.ErrorResponse(http.StatusInternalServerError, constantscore.InvalidResponseBody)
		}
		return response.SuccessResponse(http.StatusOK, responseBody, constantscore.ItemSuccessfullyObtained)
	}
	return response.SuccessResponse(http.StatusOK, responseBody, constantscore.DataNotFound)
}

// UpdateAmountBankService is responsible for the removal of a bank.
func (r *BankService) UpdateAmountBankService(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("UpdateAmountBankService", ctx, request)
	var bank models.Bank
	var responseBody []byte
	errorUnmarshal := json.Unmarshal([]byte(request.Body), &bank)
	if errorUnmarshal != nil {
		logs.LogTrackingError("UpdateAmountBankService", "JSON Unmarshal", ctx, request, errorUnmarshal)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.InvalidRequestBody)
	}
	bank.UpdatedAt = time.Now().UTC().Unix()
	errorUpdatePutRepository := r.bankRepo.UpdateAmountBankRepository(ctx, request, bank)
	if errorUpdatePutRepository != nil {
		logs.LogTrackingError("UpdateAmountBankService", "UpdateAmountBankRepository", ctx, request, errorUpdatePutRepository)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.ErrorUpdatingItem)
	}

	return response.SuccessResponse(http.StatusOK, responseBody, constantscore.ItemSuccessfullyUpdated)
}
