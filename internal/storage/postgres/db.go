package postgres

import (
	"context"
	"errors"
	"github.com/iamgafurov/journal/internal/storage"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"runtime"
)

var errNotAffected = errors.New("row is not affected")

// Db postgres struct
type db struct {
	pool *pgxpool.Pool
}

//Close Shutdown connection of postgres
func (d *db) Close() {
	d.pool.Close()
}

// New creates new connection of postgres with pgx driver
func New(connStr string) (storage.PostgresDB, error) {
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	poolConfig.AfterConnect = afterConnect
	poolConfig.MaxConns = int32(runtime.NumCPU() * 2)
	c, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return &db{pool: c}, nil
}

func afterConnect(ctx context.Context, conn *pgx.Conn) (err error) {
	return
}
