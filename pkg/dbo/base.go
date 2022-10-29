package dbo

import (
	"context"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type BaseDA struct{}

func (s BaseDA) Insert(ctx context.Context, value interface{}) (interface{}, error) {
	db, err := GetDB(ctx)
	if err != nil {
		return nil, err
	}

	return s.InsertTx(ctx, db, value)
}

func (s BaseDA) InsertTx(ctx context.Context, db *DBContext, value interface{}) (interface{}, error) {
	err := db.ResetCondition().Create(value).Error
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if ok && me.Number == 1062 {
			return 0, ErrDuplicateRecord
		}

		return nil, err
	}

	return value, nil
}

// InsertInBatches Insert records in batch. visit https://gorm.io/docs/create.html for detail
func (s BaseDA) InsertInBatches(ctx context.Context, value interface{}, batchSize int) (interface{}, error) {
	db, err := GetDB(ctx)
	if err != nil {
		return nil, err
	}

	return s.InsertInBatchesTx(ctx, db, value, batchSize)
}

// InsertInBatchesTx Insert records in batch with context. visit https://gorm.io/docs/create.html for detail
func (s BaseDA) InsertInBatchesTx(ctx context.Context, db *DBContext, value interface{}, batchSize int) (interface{}, error) {
	err := db.ResetCondition().CreateInBatches(value, batchSize).Error
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if ok && me.Number == 1062 {

			return 0, ErrDuplicateRecord
		}

		return nil, err
	}

	return value, nil
}

func (s BaseDA) Update(ctx context.Context, value interface{}) (int64, error) {
	db, err := GetDB(ctx)
	if err != nil {
		return 0, err
	}

	return s.UpdateTx(ctx, db, value)
}

func (s BaseDA) UpdateTx(ctx context.Context, db *DBContext, value interface{}) (int64, error) {
	newDB := db.ResetCondition().Save(value)
	if newDB.Error != nil {
		me, ok := newDB.Error.(*mysql.MySQLError)
		if ok && me.Number == 1062 {

			return 0, ErrDuplicateRecord
		}

		return 0, newDB.Error
	}

	return newDB.RowsAffected, nil
}

func (s BaseDA) Save(ctx context.Context, value interface{}) error {
	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	return s.SaveTx(ctx, db, value)
}

func (s BaseDA) SaveTx(ctx context.Context, db *DBContext, value interface{}) error {
	err := db.ResetCondition().Save(value).Error
	if err != nil {
		return err
	}

	return nil
}

func (s BaseDA) Get(ctx context.Context, id interface{}, value interface{}) error {
	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	return s.GetTx(ctx, db, id, value)
}

func (s BaseDA) GetTx(ctx context.Context, db *DBContext, id interface{}, value interface{}) error {
	err := db.ResetCondition().Where("id=?", id).First(value).Error
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}

	return err
}

func (s BaseDA) Query(ctx context.Context, condition Conditions, values interface{}) error {
	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	return s.QueryTx(ctx, db, condition, values)
}

func (s BaseDA) QueryTx(ctx context.Context, db *DBContext, condition Conditions, values interface{}) error {
	db.ResetCondition()

	wheres, parameters := condition.GetConditions()
	if len(wheres) > 0 {
		db.DB = db.Where(strings.Join(wheres, " and "), parameters...)
	}

	orderBy := condition.GetOrderBy()
	if orderBy != "" {
		db.DB = db.Order(orderBy)
	}

	pager := condition.GetPager()
	if pager != nil && pager.Enable() {
		// pagination
		offset, limit := pager.Offset()
		db.DB = db.Offset(offset).Limit(limit)
	}

	err := db.Find(values).Error
	if err != nil {
		return err
	}

	return nil
}

func (s BaseDA) Count(ctx context.Context, condition Conditions, values interface{}) (int, error) {
	db, err := GetDB(ctx)
	if err != nil {
		return 0, err
	}

	return s.CountTx(ctx, db, condition, values)
}

func (s BaseDA) CountTx(ctx context.Context, db *DBContext, condition Conditions, value interface{}) (int, error) {
	db.ResetCondition()

	wheres, parameters := condition.GetConditions()
	if len(wheres) > 0 {
		db.DB = db.Where(strings.Join(wheres, " and "), parameters...)
	}

	var total int64
	tableName := db.GetTableName(value)
	err := db.Table(tableName).Count(&total).Error
	if err != nil {
		return 0, err
	}

	return int(total), nil
}

func (s BaseDA) Page(ctx context.Context, condition Conditions, values interface{}) (int, error) {
	db, err := GetDB(ctx)
	if err != nil {
		return 0, err
	}

	return s.PageTx(ctx, db, condition, values)
}

func (s BaseDA) PageTx(ctx context.Context, db *DBContext, condition Conditions, values interface{}) (int, error) {
	total, err := s.CountTx(ctx, db, condition, values)
	if err != nil {
		return 0, err
	}

	err = s.QueryTx(ctx, db, condition, values)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s BaseDA) QueryRawSQL(ctx context.Context, values interface{}, sql string, parameters ...interface{}) error {
	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	return s.QueryRawSQLTx(ctx, db, values, sql, parameters...)
}

func (s BaseDA) QueryRawSQLTx(ctx context.Context, db *DBContext, values interface{}, sql string, parameters ...interface{}) error {
	err := db.ResetCondition().Raw(sql, parameters...).Find(values).Error
	if err != nil {

		return err
	}

	return nil
}
