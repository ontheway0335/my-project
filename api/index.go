package handler

import (
	"net/http"
	"one-api/common"     // 必须和 go.mod 里的 module 名称一致
	"one-api/model"
	"one-api/router"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	common.SetupLogger()
	common.Init()
	model.InitDB()
	
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	router.SetRouter(server)
	server.ServeHTTP(w, r)
}
