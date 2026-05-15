package handler

import (
	"net/http"
	"one-api/common"
	"one-api/model"
	"one-api/router"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化程序
	common.SetupLogger()
	common.Init()
	model.InitDB()
	
	// 2. 设置路由
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	router.SetRouter(server)
	
	// 3. 将 Vercel 请求转接给 Gin 框架
	server.ServeHTTP(w, r)
}
