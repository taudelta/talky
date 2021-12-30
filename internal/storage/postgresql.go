package storage

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/taudelta/talky/internal/config"
	"github.com/taudelta/talky/internal/storage/postgresql"
)

type PostgreSQLStorage struct {
	db *pgxpool.Pool
}

func NewPostgreSQLStorage(cfg *config.Config) *PostgreSQLStorage {

	poolConfig, err := pgxpool.ParseConfig(cfg.DbDSN)
	if err != nil {
		log.Println("Unable to parse DATABASE_URL", err)
		os.Exit(1)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Println("Unable to create connection pool", err)
		os.Exit(1)
	}

	return &PostgreSQLStorage{
		db: db,
	}
}

func (s *PostgreSQLStorage) Ping() error {
	err := s.db.Ping(context.Background())
	return err
}

func (s *PostgreSQLStorage) Query() *postgresql.Postgres {
	return postgresql.NewPostgres(s.db)
}

func (s *PostgreSQLStorage) QueryTx() (*postgresql.PostgresTx, error) {
	t, err := s.db.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	return postgresql.NewPostgresTx(t), nil
}
