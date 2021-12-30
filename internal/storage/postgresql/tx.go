package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type PostgresTx struct {
	Postgres
	p pgx.Tx
}

func NewPostgresTx(p pgx.Tx) *PostgresTx {
	return &PostgresTx{
		p: p,
	}
}

func (s *PostgresTx) Commit(ctx context.Context) error {
	return s.p.Commit(ctx)
}

func (s *PostgresTx) Rollback(ctx context.Context) error {
	return s.p.Rollback(ctx)
}
