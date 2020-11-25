package global

import "github.com/eden-framework/reverse-proxy/master"

var ProxyConfig = struct {
	Server *master.Master
}{
	Server: &master.Master{
		ListenAddr: ":9080",
	},
}
