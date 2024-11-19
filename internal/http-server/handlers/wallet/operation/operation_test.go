package operation_test

import (
	"context"
	"testing"
	"testproject-rest/internal/http-server/handlers/wallet/operation/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBalanceChanger_Deposit_Success(t *testing.T) {
	mockBalanceChanger := mocks.NewBalanceChanger(t)
	ctx := context.Background()
	walletUuid := uuid.New()
	amount := int64(1000)

	// Setting up expectation
	mockBalanceChanger.On("Deposit", ctx, walletUuid, amount).Return(nil)

	// Call the method
	err := mockBalanceChanger.Deposit(ctx, walletUuid, amount)

	// Assertions
	assert.NoError(t, err)
	mockBalanceChanger.AssertExpectations(t)
}

func TestBalanceChanger_Deposit_Error(t *testing.T) {
	mockBalanceChanger := mocks.NewBalanceChanger(t)
	ctx := context.Background()
	walletUuid := uuid.New()
	amount := int64(1000)
	expectedError := assert.AnError // Using a predefined error for testing

	// Setting up expectation
	mockBalanceChanger.On("Deposit", ctx, walletUuid, amount).Return(expectedError)

	// Call the method
	err := mockBalanceChanger.Deposit(ctx, walletUuid, amount)

	// Assertions
	assert.Equal(t, expectedError, err)
	mockBalanceChanger.AssertExpectations(t)
}

func TestBalanceChanger_Withdraw_Success(t *testing.T) {
	mockBalanceChanger := mocks.NewBalanceChanger(t)
	ctx := context.Background()
	walletUuid := uuid.New()
	amount := int64(500)

	// Setting up expectation
	mockBalanceChanger.On("Withdraw", ctx, walletUuid, amount).Return(nil)

	// Call the method
	err := mockBalanceChanger.Withdraw(ctx, walletUuid, amount)

	// Assertions
	assert.NoError(t, err)
	mockBalanceChanger.AssertExpectations(t)
}

func TestBalanceChanger_Withdraw_Error(t *testing.T) {
	mockBalanceChanger := mocks.NewBalanceChanger(t)
	ctx := context.Background()
	walletUuid := uuid.New()
	amount := int64(500)
	expectedError := assert.AnError // Using a predefined error for testing

	// Setting up expectation
	mockBalanceChanger.On("Withdraw", ctx, walletUuid, amount).Return(expectedError)

	// Call the method
	err := mockBalanceChanger.Withdraw(ctx, walletUuid, amount)

	// Assertions
	assert.Equal(t, expectedError, err)
	mockBalanceChanger.AssertExpectations(t)
}
