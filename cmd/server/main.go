package main

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/initialize"

	_ "github.com/DangPham112000/go-ecommerce-backend-api/cmd/swag/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Ecommerce backend API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/DangPham112000/go-ecommerce-backend-api.git

// @contact.name   Dang Pham
// @contact.url    https://google.com
// @contact.email  dangpham112000@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2024
// @schema http
func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")
}
