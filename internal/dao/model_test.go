package dao

import (
	"os"
	"testing"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

func TestMain(m *testing.M) {
	InitMySQL()

	os.Exit(m.Run())
}

func InitMySQL() {
	dboHandler, err := dbo.NewWithConfig(func(c *dbo.Config) {
		c.DBType = dbo.MySQL
		c.MaxIdleConns = 64             //dbConf.MaxIdleConns
		c.MaxOpenConns = 64             //dbConf.MaxOpenConns
		c.ConnMaxLifetime = time.Minute //dbConf.ConnMaxLifetime
		c.ConnectionString = "root:root1234@tcp(127.0.0.1:3306)/metapass?charset=utf8&parseTime=True&loc=Local"
		c.ShowLog = true
		c.LogLevel = dbo.Info
	})
	if err != nil {
		panic("create dbo failed" + err.Error())
	}
	dbo.ReplaceGlobal(dboHandler)
}
