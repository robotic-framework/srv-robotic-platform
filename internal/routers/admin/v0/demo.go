package v0

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(Demo{}))
}

//
type Demo struct {
	httpx.MethodGet
}

func (req Demo) Path() string {
	return ""
}

func (req Demo) Output(ctx context.Context) (result interface{}, err error) {
	panic("implement me")
}
