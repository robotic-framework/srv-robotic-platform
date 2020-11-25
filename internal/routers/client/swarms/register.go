package swarms

import (
	"context"
	"github.com/eden-framework/courier"
)

func init() {
	Router.Register(courier.NewRouter(RegisterSwarm{}))
}

// 注册蜂群
type RegisterSwarm struct {
}

func (req RegisterSwarm) Output(ctx context.Context) (result interface{}, err error) {
	panic("implement me")
}
