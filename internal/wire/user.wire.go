//go:build wireinject

package wire

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/controller"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/repo"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
