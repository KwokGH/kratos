package data

import (
	"context"
	"github.com/KwokGH/kratos/internal/conf"
	"github.com/KwokGH/kratos/internal/entity/de"
	"github.com/KwokGH/kratos/pkg/dbo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	redisClient *redis.Client
}

// NewData .
func NewData(cfg *conf.Data, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(logger)
	ctx := context.Background()
	// mysql gorm db
	//db, err := gorm.Open(mysql.Open(cfg.GetDatabase().GetSource()), &gorm.Config{})
	//if err != nil {
	//	l.Info("数据库初时候失败")
	//	return nil, nil, err
	//}

	if err := InitMySQL(cfg); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		l.Info("closing the data resources")
	}

	// redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		l.Info("redis连接失败")
		return nil, nil, err
	} else {
		l.Info("redis连接成功")
	}

	if err := initTable(ctx); err != nil {
		return nil, cleanup, err
	}

	return &Data{
		redisClient: rdb,
	}, cleanup, nil
}

func InitMySQL(cfg *conf.Data) error {
	dboHandler, err := dbo.NewWithConfig(func(c *dbo.Config) {
		c.DBType = dbo.MySQL
		c.ConnectionString = cfg.Database.GetSource()
		c.LogLevel = dbo.Info
	})
	if err != nil {
		return err
	}
	dbo.ReplaceGlobal(dboHandler)

	return nil
}

func initTable(ctx context.Context) error {
	err := dbo.MustGetDB(ctx).AutoMigrate(&de.User{})

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
