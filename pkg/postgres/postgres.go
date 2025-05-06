package postgres

import (
	"database/sql"
	"sync"

	"github.com/gera9/go-blog/pkg/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConn struct {
	*sql.DB
}

var (
	postgresOnce sync.Once
	postgresConn *PostgresConn
)

func GetPostgresConn(connStr string) (*PostgresConn, error) {
	var err error

	postgresOnce.Do(func() {
		var db *sql.DB

		db, err = utils.Retry(func() (*sql.DB, error) {
			db, err = sql.Open("pgx", connStr)
			if err != nil {
				return nil, err
			}

			err = db.Ping()
			if err != nil {
				return nil, err
			}

			return db, nil
		})
		if err != nil {
			return
		}

		postgresConn = &PostgresConn{db}
	})

	return postgresConn, err
}

func (p PostgresConn) Close() error {
	return p.DB.Close()
}
