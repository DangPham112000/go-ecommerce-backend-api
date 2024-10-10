package repo

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
