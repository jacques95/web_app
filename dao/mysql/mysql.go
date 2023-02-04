package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db_name"),
	)
	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
	}
	//db.SetMaxOpenConns(viper.GetInt("mysql.max_cons"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idles"))
	return
}

func Close() {
	_ = db.Close()
}
