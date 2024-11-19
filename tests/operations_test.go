package tests

import (
	"net/http"
	"testing"
	"testproject-rest/internal/lib/rest/response"
	"testproject-rest/tests/suite"

	"github.com/stretchr/testify/require"
)

func TestWalletOperation_OK(t *testing.T) {
	ctx, st := suite.New(t)

	resp, err := st.Client.R().
		SetContext(ctx).
		SetResult(&response.Response{}).
		SetBody(`{"walletId":"54573b4d-40cd-4dfa-9e60-5cf8997c6178","operationType":"DEPOSIT","amount":"10"}`).
		Post("/api/v1/wallet")

	require.NoError(t, err)

	require.Equal(t, resp.StatusCode(), http.StatusOK)

	require.JSONEq(t, string(resp.Body()), string("{\"status\":\"success\"}"))

}

func TestWalletOperation_BadRequest(t *testing.T) {
	ctx, st := suite.New(t)

	resp, err := st.Client.R().
		SetContext(ctx).
		SetResult(&response.Response{}).
		SetBody(`{"walletId":"incorrect","operationType":"DEPOSIT","amount":10}`).
		Post("/api/v1/wallet")

	require.NoError(t, err)

	require.Equal(t, resp.StatusCode(), http.StatusBadRequest)

	require.JSONEq(t, string(resp.Body()), string("{\"status\":\"error\",\"error\":\"invalid UUID length: 9\"}"))

}

func TestWalletOperation_NotFound(t *testing.T) {
	ctx, st := suite.New(t)

	resp, err := st.Client.R().
		SetContext(ctx).
		SetResult(&response.Response{}).
		SetBody(`{"walletId":"6d1edb09-e204-48cc-91bf-e7645993d073","operationType":"DEPOSIT","amount":"10"}`).
		Post("/api/v1/wallet")

	require.NoError(t, err)

	require.Equal(t, resp.StatusCode(), http.StatusNotFound)

	require.JSONEq(t, string(resp.Body()), string("{\"status\":\"error\",\"error\":\"storage.pgsql.Deposit: wallet not found\"}"))

}
