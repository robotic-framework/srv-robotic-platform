module github.com/robotic-framework/srv-robotic-platform

go 1.14

replace (
	github.com/eden-framework/context => /Users/liyiwen/Documents/golang/src/github.com/eden-framework/context
	github.com/eden-framework/reverse-proxy => /Users/liyiwen/Documents/golang/src/github.com/eden-framework/reverse-proxy
	k8s.io/client-go => k8s.io/client-go v0.18.8
)

require (
	github.com/eden-framework/apollo v0.0.1
	github.com/eden-framework/context v0.0.3
	github.com/eden-framework/courier v1.0.4
	github.com/eden-framework/eden-framework v1.1.8
	github.com/eden-framework/sqlx v0.0.1
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/henrylee2cn/ameda v1.4.5 // indirect
	github.com/henrylee2cn/erpc/v6 v6.3.4 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20200831095421-56f2037199c0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/lucas-clemente/quic-go v0.19.1 // indirect
	github.com/mmcloughlin/avo v0.0.0-20201105074841-5d2f697d268f // indirect
	github.com/eden-framework/reverse-proxy v0.0.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v0.0.5
	github.com/xtaci/kcp-go/v5 v5.6.1 // indirect
	golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/tools v0.0.0-20201116002733-ac45abd4c88c // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
