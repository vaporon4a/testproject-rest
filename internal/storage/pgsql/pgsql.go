package pgsql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrWalletNotFound = fmt.Errorf("wallet not found")
)

type Storage struct {
	db *pgxpool.Pool
}

func New(connString string) (*Storage, error) {
	const op = "storage.pgsql.New"

	db, err := pgxpool.New(context.Background(), connString)
	// db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) Deposit(ctx context.Context, walletUuid uuid.UUID, amount int64) error {
	const op = "storage.pgsql.Deposit"

	query := "UPDATE wallets SET balance = balance + @amount WHERE uuid::text = @walletUuid;"

	args := pgx.NamedArgs{
		"amount":     amount,
		"walletUuid": walletUuid,
	}
	c, err := s.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if c.RowsAffected() == 0 {
		return ErrWalletNotFound
	}

	return nil
}

func (s *Storage) Withdraw(ctx context.Context, walletUuid uuid.UUID, amount int64) error {
	const op = "storage.pgsql.Withdraw"

	query := "UPDATE wallets SET balance = balance - @amount WHERE uuid::text = @walletUuid;"

	args := pgx.NamedArgs{
		"amount":     amount,
		"walletUuid": walletUuid,
	}

	c, err := s.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if c.RowsAffected() == 0 {
		return ErrWalletNotFound
	}

	return nil
}

func (s *Storage) Balance(ctx context.Context, walletUuid uuid.UUID) (balance int64, err error) {
	const op = "storage.pgsql.Balance"

	query := "SELECT balance FROM wallets WHERE uuid::text = @walletUuid;"

	args := pgx.NamedArgs{
		"walletUuid": walletUuid,
	}

	err = s.db.QueryRow(ctx, query, args).Scan(&balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, fmt.Errorf("%s: %w", op, ErrWalletNotFound)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return balance, nil
}
