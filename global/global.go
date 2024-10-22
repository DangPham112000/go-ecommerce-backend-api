package global

import (
	"database/sql"

	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/logger"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB
	Mdbc          *sql.DB
	Rbd           *redis.Client
	KafkaProducer *kafka.Writer
)
