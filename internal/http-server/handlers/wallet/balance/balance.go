package balance

import (
	"context"
	"log/slog"
	"net/http"
	"testproject-rest/internal/lib/logger/slhelper"
	"testproject-rest/internal/lib/rest/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type Response struct {
	Balance int64  `json:"balance,omitempty"`
	Error   string `json:"error,omitempty"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.47.0 --name=BalanceShower
type BalanceShower interface {
	Balance(ctx context.Context, walletUuid uuid.UUID) (balance int64, err error)
}

func ShowBalance(log *slog.Logger, balanceShower BalanceShower) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.wallet.balance.ShowBalance"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		walletUuid, err := uuid.Parse(chi.URLParam(r, "walletId"))
		if err != nil {
			log.Error("failed to parse wallet uuid", slhelper.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err))

			return
		}

		balance, err := balanceShower.Balance(r.Context(), walletUuid)
		if err != nil {
			if err.Error() == "storage.pgsql.Balance: wallet not found" {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusInternalServerError)
			}
			log.Error("failed to get balance", slhelper.Err(err))

			render.JSON(w, r, response.Error(err))

			return
		}

		render.JSON(w, r, Response{Balance: balance})
	}
}
