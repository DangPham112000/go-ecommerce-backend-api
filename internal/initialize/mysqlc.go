package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanicC(err, "Mysql initialization error")
	global.Logger.Info("Mysql initialization success!!!")
	// global.Logger.Info(fmt.Sprintf("Mysql initialization success!!! s=%v", s))
	global.Mdbc = db

	SetPool()
	// genTableDAO()
	// migrateTables()
}

func SetPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql error :: %s", err)
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

// func genTableDAOC() {
// 	// Initiate the tables
// 	g := gen.NewGenerator(gen.Config{
// 		OutPath: "./internal/model",
// 		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
// 	})
// 	g.UseDB(global.Mdb)
// 	g.GenerateModel("go_crm_user")
// 	g.Execute()

// }

// func migrateTables() {
// 	err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})
// 	if err != nil {
// 		fmt.Println("Tables Migration error")
// 	}
// }
