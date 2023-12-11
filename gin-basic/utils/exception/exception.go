package exception

import (
	"gin-basic/common/message"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GlobalExceptionHandler 全局异常处理中间件
func GlobalExceptionHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 处理异常
			log.Println("发生了一个异常:", err)

			// 返回错误响应给客户端
			c.JSON(http.StatusInternalServerError, message.ErrorServer())
		}
	}()

	// 继续处理请求
	c.Next()
}
