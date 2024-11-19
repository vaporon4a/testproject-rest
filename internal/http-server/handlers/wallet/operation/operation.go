package operation

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
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

func UseWallet(log *slog.Logger, balanceChanger BalanceChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.wallet.operation.UseWallet"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request", slhelper.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err))
			return
		}

		log.Info("request received", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("failed to validate request", slhelper.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		var actionErr error
		switch req.OperationType {
		case "WITHDRAW":
			actionErr = balanceChanger.Withdraw(r.Context(), req.WalletId, req.Amount)
		case "DEPOSIT":
			actionErr = balanceChanger.Deposit(r.Context(), req.WalletId, req.Amount)
		default:
			log.Error("unknown operation type", slog.String("operation_type", req.OperationType))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(fmt.Errorf("unknown operation type: %s", req.OperationType)))
			return
		}

		if actionErr != nil {
			if strings.Contains(actionErr.Error(), "wallet not found") {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusInternalServerError)
			}
			log.Error("operation failed", slhelper.Err(actionErr))
			render.JSON(w, r, response.Error(actionErr))
			return
		}

		log.Info("operation success", slog.String("operation_type", req.OperationType), slog.Any("wallet", req.WalletId))
		render.JSON(w, r, response.Success())
	}
}
