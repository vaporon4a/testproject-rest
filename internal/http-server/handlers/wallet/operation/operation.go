package operation

import (
	"context"
	"log/slog"
	"net/http"
	"testproject-rest/internal/lib/logger/slhelper"
	"testproject-rest/internal/lib/rest/response"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Request struct {
	WalletId      uuid.UUID `json:"walletId" validate:"required,uuid4"`
	OperationType string    `json:"operationType" validate:"required,oneof=DEPOSIT WITHDRAW"`
	Amount        int64     `json:"amount,string" validate:"required,gte=0"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.47.0 --name=BalanceChanger
type BalanceChanger interface {
	Deposit(ctx context.Context, walletUuid uuid.UUID, amount int64) (err error)
	Withdraw(ctx context.Context, walletUuid uuid.UUID, amount int64) (err error)
}

// UseWallet returns an HTTP handler that performs either a DEPOSIT or WITHDRAW
// operation on a wallet. The request body should contain a JSON object with
// the following structure:
//
//	{
//	  "walletId": "uuid4",
//	  "operationType": "DEPOSIT" or "WITHDRAW",
//	  "amount": int64
//	}
//
// The handler first validates the request using the validator library.
// If the request is invalid, it returns a 400 error with the validation
// errors.
//
// If the request is valid, it performs the requested operation using the
// provided BalanceChanger. If the operation fails, it returns a 500 error.
//
// If the operation succeeds, it returns a 200 OK response.
func UseWallet(log *slog.Logger, balanceChanger BalanceChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.wallet.operation.Deposit"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request", slhelper.Err(err))

			render.JSON(w, r, response.Error(err))

			return
		}

		log.Info("request received", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("failed to validate request", slhelper.Err(err))

			render.JSON(w, r, response.ValidationError(validateErr))

			return
		}
		if req.OperationType == "WITHDRAW" {
			if err := balanceChanger.Withdraw(r.Context(), req.WalletId, req.Amount); err != nil {
				log.Error("failed to withdraw", slhelper.Err(err))

				render.JSON(w, r, response.Error(err))

				return
			} else {
				log.Info("withdraw success", slog.Any("wallet", req.WalletId))

				render.JSON(w, r, response.Success())
			}

		}
		if req.OperationType == "DEPOSIT" {
			if err := balanceChanger.Deposit(r.Context(), req.WalletId, req.Amount); err != nil {
				log.Error("failed to deposit", slhelper.Err(err))

				render.JSON(w, r, response.Error(err))

				return
			} else {
				log.Info("deposit success", slog.Any("wallet", req.WalletId))

				render.JSON(w, r, response.Success())
			}

		}
	}
}
