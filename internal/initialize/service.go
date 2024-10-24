package initialize

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/database"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/service"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
}
