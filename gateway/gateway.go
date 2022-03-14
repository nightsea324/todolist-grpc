package gateway

import (
	"main/gateway/controller/member"
	"main/gateway/controller/todolist"
	"main/gateway/infra/redis"
	"main/gateway/router"
)

// GatewayServer -
func GatewayServer() {
	member.New()
	todolist.New()
	redis.Init()
	router.Router()
}
