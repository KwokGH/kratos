package dbo

import (
	// mysql driver
	"context"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	globalDBO   *DBO
	globalMutex sync.Mutex
)

// DBO database operator
type DBO struct {
	db     *gorm.DB
	config *Config
}

// MustGetDB get db context otherwise panic
func MustGetDB(ctx context.Context) *DBContext {
	dbContext, err := GetDB(ctx)
	if err != nil {
		panic("get db context failed" + err.Error())
	}

	return dbContext
}

// GetDB get db context
func GetDB(ctx context.Context) (*DBContext, error) {
	dbo, err := GetGlobal()
	if err != nil {
		return nil, err
	}

	return dbo.GetDB(ctx), nil
}

// ReplaceGlobal replace global dbo instance
func ReplaceGlobal(dbo *DBO) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	globalDBO = dbo
}

// GetGlobal get global dbo
func GetGlobal() (*DBO, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if globalDBO == nil {
		dbo, err := New()
		if err != nil {
			return nil, err
		}

		globalDBO = dbo
	}

	return globalDBO, nil
}

// New create new database operator
func New(options ...Option) (*DBO, error) {
	return NewWithConfig(options...)
}

// NewWithConfig create new database operator
func NewWithConfig(options ...Option) (*DBO, error) {
	// init config with default values
	config := getDefaultConfig()

	for _, option := range options {
		option(config)
	}

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,   // 慢 SQL 阈值
	// 		LogLevel:                  logger.Silent, // 日志级别
	// 		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
	// 		Colorful:                  false,         // 禁用彩色打印
	// 	},
	// )

	var db *gorm.DB
	var err error
	switch config.DBType {
	case MySQL:
		db, err = gorm.Open(mysql.New(mysql.Config{
			DriverName: config.DBType.DriverName(),
			DSN:        config.ConnectionString,
		}), &gorm.Config{
			QueryFields: true,
			Logger:      logger.Default.LogMode(logger.Info),
		})
	default:
		panic("unsupported database type" + config.DBType.String())
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	if config.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}

	if config.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}

	if config.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	}

	if config.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	}

	return &DBO{db, config}, nil
}

func (s DBO) GetDB(ctx context.Context) *DBContext {
	ctxDB := &DBContext{DB: s.db.Session(&gorm.Session{
		Context:     ctx,
		NewDB:       true,
		QueryFields: true,
	})}

	ctxDB.Logger = logger.New(ctxDB, logger.Config{
		LogLevel:                  s.config.LogLevel.GormLogLevel(),
		SlowThreshold:             s.config.SlowThreshold,
		IgnoreRecordNotFoundError: false,
		Colorful:                  false,
	})

	return ctxDB
}
