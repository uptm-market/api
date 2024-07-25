package infradb

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	// PostgreSQL Driver (integrado com New Relic)
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

var (
	DB   *sql.DB
	dbMu sync.Mutex

	// DSN string de conexão
	DSN             string
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
)

// Tx interface comum para estruturas com transação e sem transação.
type Tx interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type DBConn interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Ping() error
	Tx
}

// newDB cria e abre uma nova conexão com o banco de dados.
func newDB(ds string) (*sql.DB, error) {
	db, err := sql.Open("nrpostgres", ds)
	if err != nil {
		return nil, err
	}
	err = checkDBIsUp(db)
	return db, err
}

func checkDBIsUp(db *sql.DB) error {
	const maxRetries = 3
	var (
		up      bool
		retries int
		err     error
	)
	parentCtx := context.Background()
	log.Println("teste")
	for !up {
		err = func() error {
			ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
			defer cancel()
			if err := db.PingContext(ctx); err != nil {
				return err
			}
			up = true
			return nil
		}()
		if err != nil {
			if retries < maxRetries-1 {
				retries++
				time.Sleep(1 * time.Second)
			} else {
				up = true
			}
		}
	}
	return err
}

// Get retorna uma conexão com o banco de dados (abre se necessário).
func Get() *sql.DB {

	dbMu.Lock()
	defer dbMu.Unlock()
	if DB == nil {
		var err error
		DB, err = newDB("postgresql://postgres:KGzQGxTwRlIYHxKdSsCPKhMiBbbmiRhM@monorail.proxy.rlwy.net:46803/railway")
		if err != nil {
			log.Fatalln("get db:", err.Error())
		}
	}
	return DB
}

// GetTx obtém uma transação mesmo se tx for nulo.
func GetTx(tx *sql.Tx) Tx {
	if tx == nil {
		return Get()
	}
	return tx
}
