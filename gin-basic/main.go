package main

import (
	"gin-basic/config"
	"gin-basic/utils/exception"
	"gin-basic/utils/mysql"
	"gin-basic/utils/redis"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func init() {
	mysql.Init()
	redis.Init()
}

func main() {
	r := gin.Default()
	r.Use(exception.GlobalExceptionHandler)

	// 启动服务
	portStr := ":" + strconv.Itoa(config.Config.Server.Port)
	if err := r.Run(portStr); err != nil {
		log.Error(err)
	}
}
