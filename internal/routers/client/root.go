package client

import (
	"github.com/eden-framework/courier"
	"github.com/robotic-framework/srv-robotic-platform/internal/routers/client/swarms"
)

var Router = courier.NewRouter(Group{})

func init() {
	Router.Register(swarms.Router)
}

type Group struct {
	courier.EmptyOperator
}
