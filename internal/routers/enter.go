package routers

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/routers/manage"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
