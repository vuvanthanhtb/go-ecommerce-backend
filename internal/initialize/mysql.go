package initialize

import (
	"fmt"
	"time"

	"github.com/vuvanthanhtb/go-ecommerce-backend/global"
	po "github.com/vuvanthanhtb/go-ecommerce-backend/internal/persistence_object"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.MySQL

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Initializing MySQL error")
	global.Logger.Info("Initializing MySQL successfully")
	global.Mdb = db

	SetPool()
	migrateTables()
}

func SetPool() {
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("mysql err: ", err)
	}

	m := global.Config.MySQL
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Println("migrating tables error: ", err)
	}
}
