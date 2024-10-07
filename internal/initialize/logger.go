package initialize

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
