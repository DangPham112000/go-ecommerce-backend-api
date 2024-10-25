package initialize

import (
	"fmt"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
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
	return r
}
