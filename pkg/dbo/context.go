package dbo

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// DBContext db with context
type DBContext struct {
	*gorm.DB
}

// Print print sql log
func (s *DBContext) Printf(format string, v ...interface{}) {
	switch len(v) {
	case 4:
		// v[4]: [fileWithLineNum(), duration, rowAffected, sql]
		// example: ["/home/test/code.go:27",16.845856,1,"INSERT INTO `test_table` (`id`,`name`) VALUES ('1234','bad')"]
		logx.Info(s.Statement.Context, v[3].(string),
			logx.Field("logType", "sql"),
			logx.Field("lineNum", v[0].(string)),
			logx.Field("rowsAffected", v[2]),
			logx.Field("duration", v[1].(float64)))
	case 5:
		// v[5]: [fileWithLineNum(), "SLOW SQL >= 1µs", duration, rowAffected, sql]
		// example: ["/home/test/code.go:27","SLOW SQL >= 1µs",16.845856,1,"INSERT INTO `test_table` (`id`,`name`) VALUES ('1234','bad')"]
		logx.Info(s.Statement.Context, v[4].(string),
			logx.Field("logType", "sql"),
			logx.Field("lineNum", v[0].(string)),
			logx.Field("rowsAffected", v[3]),
			logx.Field("duration", v[2].(float64)),
			logx.Field("extra", v[1]))
	default:
		logx.Info(s.DB.Statement.Context, "invalid sql log",
			logx.Field("logType", "sql"),
			logx.Field("format", format),
			logx.Field("args", v))
	}
}

// GetTableName get database table name of value
func (s *DBContext) GetTableName(value interface{}) string {
	stmt := &gorm.Statement{DB: s.DB}
	err := stmt.Parse(value)
	if err != nil {
		return ""
	}

	return stmt.Schema.Table
}

// ResetCondition reset session query conditions
func (s *DBContext) ResetCondition() *DBContext {
	s.DB = s.DB.Session(&gorm.Session{NewDB: true})
	return s
}
