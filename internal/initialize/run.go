package initialize

import (
	"fmt"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading mysql config", m.Host, m.Dbname)
	InitLogger()
	global.Logger.Info("Config ok!", zap.String("ok", "success"))
	InitMysql()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()
	InitKafka()
	r := InitRouter()
	r.Run(":8002")
}
