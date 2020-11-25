package admin

import (
	"github.com/eden-framework/context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/swagger"

	v0 "github.com/robotic-framework/srv-robotic-platform/internal/routers/admin/v0"
)

var Router = courier.NewRouter(RootRouter{})

func init() {
	Router.Register(v0.Router)
	if !context.IsOnline() {
		Router.Register(swagger.SwaggerRouter)
	}
}

type RootRouter struct {
	courier.EmptyOperator
}

func (RootRouter) Path() string {
	return "/robotic-platform"
}
