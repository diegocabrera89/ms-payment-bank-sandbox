package repository

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/constantsmicro"
	"github.com/diegocabrera89/ms-payment-bank-sandbox/models"
	"github.com/diegocabrera89/ms-payment-core/dynamodbcore"
	"github.com/diegocabrera89/ms-payment-core/helpers"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"os"
)

// BankRepositoryImpl implements the BankRepository interface of the ms-payment-core package.
type BankRepositoryImpl struct {
	CoreRepository dynamodbcore.CoreRepository
}

// NewBankRepository create a new BankRepository instance.
func NewBankRepository() *BankRepositoryImpl {
	bankTable := os.Getenv(constantsmicro.BankTableName)
	region := os.Getenv(constantsmicro.Region)

	coreRepository, _ := dynamodbcore.NewDynamoDBRepository(bankTable, region)

	return &BankRepositoryImpl{
		CoreRepository: coreRepository,
	}
}

// CreateBankRepository put item in DynamoDB.
func (r *BankRepositoryImpl) CreateBankRepository(ctx context.Context, request events.APIGatewayProxyRequest, bank models.Bank) (models.Bank, error) {
	logs.LogTrackingInfo("CreateBankRepository", ctx, request)
	item, errorMarshallItem := helpers.MarshallItem(bank)
	if errorMarshallItem != nil {
		logs.LogTrackingError("CreateBankRepository", "MarshallItem", ctx, request, errorMarshallItem)
		return models.Bank{}, errorMarshallItem
	}

	errorPutItemCore := r.CoreRepository.PutItemCore(ctx, request, item)
	if errorPutItemCore != nil {
		return models.Bank{}, errorPutItemCore
	}
	return bank, nil
}

// GetBankByBankIDRepository put item in DynamoDB.
func (r *BankRepositoryImpl) GetBankByBankIDRepository(ctx context.Context, request events.APIGatewayProxyRequest, fieldValueFilterByID string) (models.Bank, error) {
	logs.LogTrackingInfo("GetBankByBankIDRepository", ctx, request)
	var bank models.Bank
	responseGetBankById, errorGetBankByBankID := r.CoreRepository.GetItemCore(ctx, request, constantsmicro.BankID, fieldValueFilterByID)
	if errorGetBankByBankID != nil {
		logs.LogTrackingError("GetBankByBankIDRepository", "GetItemCore", ctx, request, errorGetBankByBankID)
		return models.Bank{}, errorGetBankByBankID
	}
	errUnmarshalMap := helpers.UnmarshalMapToType(responseGetBankById.Item, &bank)
	if errUnmarshalMap != nil {
		logs.LogTrackingError("GetBankByBankIDRepository", "UnmarshalMap", ctx, request, errUnmarshalMap)
		return models.Bank{}, errUnmarshalMap
	}
	return bank, nil
}

// UpdateAmountBankRepository update item in DynamoDB.
func (r *BankRepositoryImpl) UpdateAmountBankRepository(ctx context.Context, request events.APIGatewayProxyRequest, bank models.Bank) error {
	logs.LogTrackingInfo("UpdateAmountBankRepository", ctx, request)
	skipFields := []string{constantsmicro.BankID, constantsmicro.AccountID, constantsmicro.CreatedAt}
	errorUpdatePut := r.CoreRepository.UpdateItemCore(ctx, request, bank, constantsmicro.BankID, bank.BankID, skipFields)
	if errorUpdatePut != nil {
		logs.LogTrackingError("UpdateAmountBankRepository", "UpdateItemCore", ctx, request, errorUpdatePut)
		return errorUpdatePut
	}
	return nil
}
