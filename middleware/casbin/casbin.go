package casbin

import (
	"github.com/gin-gonic/gin"
	mlog "mall/log"
	"mall/model"
	"mall/util"
	"net/http"
)

var log *mlog.Log

func CasbinMiddleware(c *gin.Context) {
	user := c.GetString("user") // 假设用户信息已经通过 auth.ParseToken 解析
	path := c.Request.URL.Path
	method := c.Request.Method

	allowed, err := enforcer.Enforce(user, path, method)
	if err != nil {
		log.Info("direct check auth fail:" + err.Error())
		util.Response(c, model.FORBIDDEN, "check auth fail")
		c.Abort()
		return
	}

	if !allowed {
		util.Response(c, model.FORBIDDEN, "not allowed")
		c.Abort()
		return
	}

	c.Next()
}
