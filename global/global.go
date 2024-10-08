package global

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/logger"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rbd    *redis.Client
)
