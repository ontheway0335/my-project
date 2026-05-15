package handler

import (
	"net/http"
	"one-api/common"
	"one-api/model"
	"one-api/router"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 初始化配置
	common.SetupLogger()
	common.Init()
	model.InitDB()
	
	// 设置路由
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	router.SetRouter(server)
	
	// 响应请求
	server.ServeHTTP(w, r)
}
