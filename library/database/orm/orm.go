package orm

import (
	"github.com/Guiyunweb/go-web-template/library/log"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

type Datasource struct {
	Database        string `toml:"database"`
	DSN             string `toml:"dsn"`
	ShowSql         bool   `toml:"show-sql"`
	MaxIdleConn     int    `toml:"maxIdleConn"`
	MaxOpenConn     int    `toml:"maxOpenConn"`
	ValidationQuery string `toml:"validationQuery"`
}

var DB *xorm.Engine

func NewMySQL(c *Datasource) {
	engine, err := xorm.NewEngine(c.Database, c.DSN)
	if err != nil {
		log.Panic(err.Error())
	}

	// 是否启动日志
	engine.ShowSQL(c.ShowSql)
	// 设置空闲数大小
	engine.SetMaxIdleConns(c.MaxIdleConn)
	//最大打开连接数
	engine.SetMaxOpenConns(c.MaxOpenConn)
	//最大生存时间
	engine.SetConnMaxLifetime(time.Second * 20)
	//测试数据库连接
	if _, err := engine.Query(c.ValidationQuery); err != nil {
		log.Error(err.Error())
		log.Panic("数据库连接失败")
	}

	// 测试连接是否成功
	DB = engine
}
