package operation_test

import (
	"testing"
	"testproject-rest/internal/http-server/handlers/wallet/operation/mocks"
)

func TestOperation(t *testing.T) {
	cases := []struct {
		name          string
		uuid          string
		operationType string
		amount        int64
		respError     string
		mockError     error
	}{
		{
			name:          "Success",
			uuid:          "7901fefc-5330-47ec-a3ba-192de9c636a1",
			operationType: "DEPOSIT",
			amount:        1000,
		},
		{
			name:          "Wrong UUID",
			uuid:          "wrong-uuid",
			operationType: "DEPOSIT",
			amount:        1000,
			respError:     "invalid uuid",
		},
		{
			name:          "Wrong operation type",
			uuid:          "7901fefc-5330-47ec-a3ba-192de9c636a1",
			operationType: "WRONG",
			amount:        1000,
			respError:     "invalid operation type",
		},
		{
			name:          "Negative amount",
			uuid:          "7901fefc-5330-47ec-a3ba-192de9c636a1",
			operationType: "DEPOSIT",
			amount:        -1000,
			respError:     "invalid amount",
		},
		{
			name:          "Negative amount",
			uuid:          "7901fefc-5330-47ec-a3ba-192de9c636a1",
			operationType: "WITHDRAW",
			amount:        -1000,
			respError:     "invalid amount",
		},
		{
			name:          "Zero amount",
			uuid:          "7901fefc-5330-47ec-a3ba-192de9c636a1",
			operationType: "DEPOSIT",
			amount:        0,
			respError:     "invalid amount",
		},
		{
			name:          "Zero amount",
			uuid:          "7901fefc-5330-47ec-a3ba-192de9c636a1",
			operationType: "WITHDRAW",
			amount:        0,
			respError:     "invalid amount",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			balanceChangerMock := mocks.NewBalanceChanger(t)

			if tc.mockError != nil || tc.respError != "" {
				balanceChangerMock.On("Deposit", tc.uuid, tc.amount)
			}

		})
	}
}
