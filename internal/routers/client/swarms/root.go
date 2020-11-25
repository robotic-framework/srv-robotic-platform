package swarms

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(Group{})

type Group struct {
	courier.EmptyOperator
}
