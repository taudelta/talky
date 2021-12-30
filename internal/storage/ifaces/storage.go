package ifaces

import (
	sq "github.com/Masterminds/squirrel"
)

type StorageQuery interface {
	Create(query sq.InsertBuilder, model ...interface{}) error
	Get(query sq.SelectBuilder, model ...interface{}) error
	GetAll(query sq.SelectBuilder, scannerFunc func() []interface{}) error
	Update(query sq.UpdateBuilder, model ...interface{}) error
}
