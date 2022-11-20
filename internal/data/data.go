package data

import (
	"context"
	"github.com/KwokGH/kratos/internal/conf"
	"github.com/KwokGH/kratos/internal/entity/de"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db          *gorm.DB
	redisClient *redis.Client
}

// NewData .
func NewData(cfg *conf.Data, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(logger)

	// mysql gorm db
	db, err := gorm.Open(mysql.Open(cfg.GetDatabase().GetSource()), &gorm.Config{})
	if err != nil {
		l.Info("数据库初时候失败")
		return nil, nil, err
	}

	cleanup := func() {
		l.Info("closing the data resources")
	}

	// redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		l.Info("redis连接失败")
		return nil, nil, err
	} else {
		l.Info("redis连接成功")
	}

	if err := initTable(db); err != nil {
		return nil, cleanup, err
	}

	return &Data{
		db:          db,
		redisClient: rdb,
	}, cleanup, nil
}

func initTable(db *gorm.DB) error {
	err := db.AutoMigrate(
		&de.User{})

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
