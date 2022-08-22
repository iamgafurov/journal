package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/iamgafurov/journal/internal/storage"
)

type db struct {
	pool *sql.DB
}

func New(ctx context.Context, connStr string) (mssqlDb storage.MSSQLDB, err error) {
	connString := fmt.Sprintf(connStr)
	pool, err := sql.Open("mssql", connString)
	if err != nil {
		return
	}
	conn, err := pool.Conn(ctx)
	if err != nil {
		return
	}
	defer conn.Close()
	err = conn.PingContext(ctx)
	if err != nil {
		return
	}

	return &db{pool: pool}, nil
}

func (d *db) Close() {
	d.pool.Close()
}
