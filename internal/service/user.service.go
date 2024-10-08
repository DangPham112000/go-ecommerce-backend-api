package service

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/repo"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/response"
)

// Interface
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 1. Check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExist
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
