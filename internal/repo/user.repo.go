package repo

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	user, err := ur.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}
	return user.UsrID != 0
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
