package postgresql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	p *pgxpool.Pool
}

func NewPostgres(p *pgxpool.Pool) *Postgres {
	return &Postgres{
		p: p,
	}
}

func (s *Postgres) Create(query sq.InsertBuilder, model ...interface{}) error {
	q, args, err := query.ToSql()
	if err != nil {
		return err
	}
	row := s.p.QueryRow(context.Background(), q, args...)
	if err := row.Scan(model...); err != nil {
		return err
	}
	return nil
}

func (s *Postgres) Get(query sq.SelectBuilder, model ...interface{}) error {
	q, args, err := query.ToSql()
	if err != nil {
		return err
	}
	row := s.p.QueryRow(context.Background(), q, args...)
	if err := row.Scan(model...); err != nil {
		return err
	}
	return nil
}

func (s *Postgres) GetAll(query sq.SelectBuilder, scannerFunc func() []interface{}) error {
	q, args, err := query.ToSql()
	if err != nil {
		return err
	}
	rows, err := s.p.Query(context.Background(), q, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(scannerFunc()...); err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (s *Postgres) Update(query sq.UpdateBuilder, model ...interface{}) error {
	q, args, err := query.ToSql()
	if err != nil {
		return err
	}
	if len(model) == 0 {
		res, err := s.p.Exec(context.Background(), q, args...)
		if err != nil {
			return err
		}
		if res.RowsAffected() == 0 {
			return sql.ErrNoRows
		}
	} else {
		row := s.p.QueryRow(context.Background(), q, args...)
		if err := row.Scan(model...); err != nil {
			return err
		}
	}
	return nil
}
