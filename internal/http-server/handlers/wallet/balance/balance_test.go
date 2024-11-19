package balance_test

import (
	"context"
	"testing"
	"testproject-rest/internal/http-server/handlers/wallet/balance/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBalanceShower_Balance(t *testing.T) {
	// Create a new instance of the mock
	mockBalanceShower := mocks.NewBalanceShower(t)

	// Prepare test data
	ctx := context.Background()
	walletUuid := uuid.New()
	expectedBalance := int64(1000)
	expectedError := error(nil)

	// Set up the expectation
	mockBalanceShower.On("Balance", ctx, walletUuid).Return(expectedBalance, expectedError)

	// Call the method under test
	balance, err := mockBalanceShower.Balance(ctx, walletUuid)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedBalance, balance)

	// Ensure that the expectations were met
	mockBalanceShower.AssertExpectations(t)
}

func TestBalanceShower_Balance_Error(t *testing.T) {
	// Create a new instance of the mock
	mockBalanceShower := mocks.NewBalanceShower(t)

	// Prepare test data
	ctx := context.Background()
	walletUuid := uuid.New()
	expectedError := assert.AnError // Using a predefined error for testing

	// Set up the expectation with an error return
	mockBalanceShower.On("Balance", ctx, walletUuid).Return(int64(0), expectedError)

	// Call the method under test
	balance, err := mockBalanceShower.Balance(ctx, walletUuid)

	// Assert the results
	assert.Equal(t, expectedError, err)
	assert.Equal(t, int64(0), balance)

	// Ensure that the expectations were met
	mockBalanceShower.AssertExpectations(t)
}
