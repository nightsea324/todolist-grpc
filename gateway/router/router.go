package router

import (
	"main/gateway/controller/member"
	"main/gateway/controller/todolist"
	"main/gateway/middleware"

	"github.com/gin-gonic/gin"
)

// Router -
func Router() {
	router := gin.Default()
	// 待辦事項
	t := router.Group("api/todolist")
	{
		// 新增待辦事項
		t.POST("/", middleware.Auth(), todolist.Create)

		// 刪除待辦事項
		t.DELETE("/:id", middleware.Auth(), todolist.Delete)

		// 完成待辦事項
		t.PUT("/:id", middleware.Auth(), todolist.Update)

		// 取得待辦事項
		t.GET("/", middleware.Auth(), todolist.Get)
	}
	// 會員
	m := router.Group("api/member")
	{
		// 會員註冊
		m.POST("/register", member.Register)

		// 會員登入
		m.POST("/login", member.SignIn)
	}
	router.Run(":80")
}
