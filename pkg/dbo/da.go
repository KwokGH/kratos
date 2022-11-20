package dbo

import (
	"context"
)

// DataAccesser data access contract
type DataAccesser interface {
	Inserter
	Updater
	Saver
	Geter
	Querier
}

type Inserter interface {
	Insert(context.Context, interface{}) (interface{}, error)
	InsertTx(context.Context, *DBContext, interface{}) (interface{}, error)
	InsertInBatches(context.Context, interface{}, int) (interface{}, error)
	InsertInBatchesTx(context.Context, *DBContext, interface{}, int) (interface{}, error)
}

type Updater interface {
	Update(context.Context, interface{}) (int64, error)
	UpdateTx(context.Context, *DBContext, interface{}) (int64, error)
}

type Saver interface {
	Save(context.Context, interface{}) error
	SaveTx(context.Context, *DBContext, interface{}) error
}

type Geter interface {
	Get(ctx context.Context, id interface{}, value interface{}) error
	GetTx(ctx context.Context, tx *DBContext, id interface{}, value interface{}) error
}

type Querier interface {
	First(ctx context.Context, condition Conditions, value interface{}) error
	FindTx(ctx context.Context, tx *DBContext, condition Conditions, value interface{}) error
	Query(ctx context.Context, condition Conditions, values interface{}) error
	QueryTx(ctx context.Context, tx *DBContext, condition Conditions, values interface{}) error
	Count(ctx context.Context, condition Conditions, entity interface{}) (int, error)
	CountTx(ctx context.Context, tx *DBContext, condition Conditions, entity interface{}) (int, error)
	Page(ctx context.Context, condition Conditions, values interface{}) (int, error)
	PageTx(ctx context.Context, tx *DBContext, condition Conditions, values interface{}) (int, error)
	QueryRawSQL(context.Context, interface{}, string, ...interface{}) error
	QueryRawSQLTx(context.Context, *DBContext, interface{}, string, ...interface{}) error
}

type Conditions interface {
	GetConditions() ([]string, []interface{})
	GetPager() *Pager
	GetOrderBy() string
}
