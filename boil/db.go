package boil

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXPoolContextExecutor struct {
	*pgxpool.Pool
}

func (p *PGXPoolContextExecutor) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	return p.Pool.Exec(context.Background(), query, args...)
}
func (p *PGXPoolContextExecutor) Query(query string, args ...interface{}) (pgx.Rows, error) {
	return p.Pool.Query(context.Background(), query, args...)
}
func (p *PGXPoolContextExecutor) QueryRow(query string, args ...interface{}) pgx.Row {
	return p.Pool.QueryRow(context.Background(), query, args...)
}
func (p *PGXPoolContextExecutor) ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return p.Pool.Exec(ctx, query, args...)
}
func (p *PGXPoolContextExecutor) QueryContext(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return p.Pool.Query(ctx, query, args...)
}
func (p *PGXPoolContextExecutor) QueryRowContext(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return p.Pool.QueryRow(ctx, query, args...)
}

var s ContextExecutor = &PGXPoolContextExecutor{}

// Executor can perform SQL queries.
type Executor interface {
	Exec(query string, args ...interface{}) (pgconn.CommandTag, error)
	Query(query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(query string, args ...interface{}) pgx.Row
}

// ContextExecutor can perform SQL queries with context
type ContextExecutor interface {
	Executor

	ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) pgx.Row
}

// Transactor can commit and rollback, on top of being able to execute queries.
type Transactor interface {
	Commit(ctx context.Context) error
	Rollback() error

	Executor
}

// Beginner begins transactions.
type Beginner interface {
	Begin() (pgx.Tx, error)
}
type PGXPoolTransactor struct {
	pgx.Tx
}

func (p *PGXPoolTransactor) Rollback() error {
	return p.Tx.Rollback(context.Background())
}
func (p *PGXPoolTransactor) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	return p.Tx.Exec(context.Background(), query, args...)
}
func (p *PGXPoolTransactor) Query(query string, args ...interface{}) (pgx.Rows, error) {
	return p.Tx.Query(context.Background(), query, args...)
}
func (p *PGXPoolTransactor) QueryRow(query string, args ...interface{}) pgx.Row {
	return p.Tx.QueryRow(context.Background(), query, args...)
}

// Begin a transaction with the current global database handle.
func Begin() (Transactor, error) {
	creator, ok := currentDB.(Beginner)
	if !ok {
		panic("database does not support transactions")
	}

	tx, err := creator.Begin()
	if err != nil {
		return nil, err
	}
	return &PGXPoolTransactor{tx}, nil
}

// ContextTransactor can commit and rollback, on top of being able to execute
// context-aware queries.
type ContextTransactor interface {
	Commit(ctx context.Context) error
	Rollback() error

	ContextExecutor
}
type PGXPoolContextTransactor struct {
	pgx.Tx
}

func (p *PGXPoolContextTransactor) ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return p.Tx.Exec(ctx, query, args...)
}
func (p *PGXPoolContextTransactor) QueryContext(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return p.Tx.Query(ctx, query, args...)
}
func (p *PGXPoolContextTransactor) QueryRowContext(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return p.Tx.QueryRow(ctx, query, args...)
}

// ContextBeginner allows creation of context aware transactions with options.
type ContextBeginner interface {
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

// BeginTx begins a transaction with the current global database handle.
func BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error) {
	creator, ok := currentDB.(ContextBeginner)
	if !ok {
		panic("database does not support context-aware transactions")
	}
	tx, err := creator.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &PGXPoolContextTransactor{tx}, nil

}
