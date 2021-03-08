package postgres

import (
	"time"

	"manger/pkg/misc/config"

	db "github.com/jinzhu/gorm"

	// go-lint
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	driverName = "postgres"
)

// New 创建数据库连接
func New(conf config.Postgres) (*db.DB, error) {
	DB, err := db.Open(driverName, conf.String())
	if err != nil {
		return nil, err
	}

	DB.DB().SetMaxIdleConns(conf.MaxIdle)
	DB.DB().SetMaxOpenConns(conf.MaxOpen)
	DB.DB().SetConnMaxLifetime(3 * time.Second)
	DB.SingularTable(true)

	DB = DB.LogMode(true)

	// 关闭 created_at updated_at deleted_at
	DB = DB.Unscoped()

	return DB, nil
}
