package routers

import (
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/routers/manage"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApplication = new(RouterGroup)
