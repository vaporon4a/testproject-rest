package tests

import (
	"net/http"
	"testing"
	"testproject-rest/internal/lib/rest/response"
	"testproject-rest/tests/suite"

	"github.com/stretchr/testify/require"
)

func TestWalletBalance_OK(t *testing.T) {
	ctx, st := suite.New(t)

	resp, err := st.Client.R().
		SetContext(ctx).
		SetResult(&response.Response{}).
		SetPathParam("wallet_id", "6d1edb09-e204-48cc-91bf-e7645993d072").
		Get("/api/v1/wallets/{wallet_id}")

	require.NoError(t, err)

	require.Equal(t, resp.StatusCode(), http.StatusOK)

	require.JSONEq(t, string(resp.Body()), string("{\"balance\":500000}"))

}

func TestWalletBalance_NotFound(t *testing.T) {
	ctx, st := suite.New(t)

	resp, err := st.Client.R().
		SetContext(ctx).
		SetResult(&response.Response{}).
		SetPathParam("wallet_id", "6d1edb09-e204-48cc-91bf-e7645993d073").
		Get("/api/v1/wallets/{wallet_id}")

	require.NoError(t, err)

	require.Equal(t, resp.StatusCode(), http.StatusNotFound)

	require.JSONEq(t, string(resp.Body()), string("{\"status\":\"error\",\"error\":\"storage.pgsql.Balance: wallet not found\"}"))

}

func TestWalletBalance_BadRequest(t *testing.T) {
	ctx, st := suite.New(t)

	resp, err := st.Client.R().
		SetContext(ctx).
		SetResult(&response.Response{}).
		SetPathParam("wallet_id", "incorrect").
		Get("/api/v1/wallets/{wallet_id}")

	require.NoError(t, err)

	require.Equal(t, resp.StatusCode(), http.StatusBadRequest)

	require.JSONEq(t, string(resp.Body()), string("{\"status\":\"error\",\"error\":\"invalid UUID length: 9\"}"))

}
