package handler

import (
	"net/http"
	"one-api/common"
	"one-api/model"
	"one-api/router"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化系统配置和数据库
	common.SetupLogger()
	common.Init()
	model.InitDB()
	
	// 2. 设置 Gin 运行模式（生产模式）
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	
	// 3. 关联路由
	router.SetRouter(server)
	
	// 4. 将 Vercel 的请求转发给 Gin 处理器
	server.ServeHTTP(w, r)
}
