package infradb

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

var (
	DB   *sql.DB
	dbMu sync.Mutex

	DSN             string
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
)

// Tx interface comum para estruturas com transação e sem transação.
type Tx interface {
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
	db, err := sql.Open("postgres", "user=user password=password dbname=database host=postgres port=5451 sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir conexão com o banco de dados: %v", err)
	}
	fmt.Println(err)
	// Configurações do pool de conexão
	db.SetConnMaxLifetime(ConnMaxLifetime)
	db.SetConnMaxIdleTime(ConnMaxIdleTime)
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	return db, nil
}

func checkDBIsUp(db *sql.DB) error {
	const maxRetries = 3
	var (
		up      bool
		retries int
		err     error
	)

	parentCtx := context.Background()

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

// LoadConfig lê as configurações do banco de dados a partir das variáveis de ambiente.
func LoadConfig() {
	DSN = os.Getenv("DB_DSN")
	fmt.Println(DSN)
	if DSN == "" {
		log.Fatalln("get db: DSN não informado")
	}

	// Configurações de outras variáveis podem ser lidas da mesma forma
}

// Load inicializa a conexão com o banco de dados.
func Load() error {
	LoadConfig()

	dbMu.Lock()
	defer dbMu.Unlock()

	if DB == nil {
		var err error
		DB, err = newDB(DSN)
		if err != nil {
			log.Fatalf("get db: falha ao iniciar a conexão com o banco de dados: %v", err)
			return err
		}
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return nil
}

// Close fecha a conexão com o banco de dados.
func Close() {
	dbMu.Lock()
	defer dbMu.Unlock()

	if DB != nil {
		DB.Close()
		fmt.Println("Conexão com o banco de dados fechada.")
	} else {
		fmt.Println("Conexão com o banco de dados já está fechada.")
	}
}

func QueryWithContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	dbMu.Lock()
	defer dbMu.Unlock()

	if DB == nil {
		return nil, fmt.Errorf("o banco de dados não está inicializado")
	}

	// Cria um contexto com um timeout de 5 segundos (você pode ajustar conforme necessário)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Executa a consulta usando o contexto e os argumentos fornecidos
	rows, err := DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("falha ao executar consulta no banco de dados: %v", err)
	}
	defer rows.Close()

	// Processar os resultados da consulta, se necessário

	return rows, nil
}
