package initialize

import (
	"fmt"
	"time"

	"github.com/vuvanthanhtb/go-ecommerce-backend/global"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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
	// migrateTables()
	// genTableDAO()
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
		// &po.User{},
		// &po.Role{},
		&model.GoCrmUserV2{},
	)

	if err != nil {
		fmt.Println("migrating tables error: ", err)
	}
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	g.GenerateModel("go_crm_user")
	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
