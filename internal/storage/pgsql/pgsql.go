package pgsql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connString string) (*Storage, error) {
	const op = "storage.pgsql.New"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Deposit(ctx context.Context, walletUuid uuid.UUID, amount int64) error {
	const op = "storage.pgsql.Deposit"

	stmt, err := s.db.PrepareContext(ctx, "UPDATE wallets SET balance = balance + ? WHERE uuid = ?;")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := stmt.ExecContext(ctx, amount, walletUuid); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) Withdraw(ctx context.Context, walletUuid uuid.UUID, amount int64) error {
	const op = "storage.pgsql.Withdraw"

	stmt, err := s.db.PrepareContext(ctx, "UPDATE wallets SET balance = balance - ? WHERE uuid = ?;")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := stmt.ExecContext(ctx, amount, walletUuid); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) Balance(ctx context.Context, walletUuid uuid.UUID) (balance int64, err error) {
	const op = "storage.pgsql.Balance"

	stmt, err := s.db.PrepareContext(ctx, "SELECT balance FROM wallets WHERE uuid = ?;")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	if err := stmt.QueryRowContext(ctx, walletUuid).Scan(&balance); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return balance, nil
}
