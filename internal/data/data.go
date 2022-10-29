package data

import (
	"context"
	"fmt"
	"github.com/KwokGH/kratos/internal/conf"
	"github.com/KwokGH/kratos/pkg/dbo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	redisClient *redis.Client
}

// NewData .
func NewData(cfg *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	initMysql(cfg)

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		fmt.Println("redis连接失败")
		panic(err)
	} else {
		fmt.Println("redis连接成功")
	}

	return &Data{
		redisClient: rdb,
	}, cleanup, nil
}

func initMysql(cfg *conf.Data) {
	dboHandler, err := dbo.NewWithConfig(func(c *dbo.Config) {
		c.DBType = dbo.DBType(cfg.Database.Driver)
		c.MaxIdleConns = 64             //dbConf.MaxIdleConns
		c.MaxOpenConns = 64             //dbConf.MaxOpenConns
		c.ConnMaxLifetime = time.Minute //dbConf.ConnMaxLifetime
		c.ConnectionString = cfg.Database.Source
		c.LogLevel = dbo.Info
	})
	if err != nil {
		panic("create dbo failed" + err.Error())
	}

	dbo.ReplaceGlobal(dboHandler)
}
