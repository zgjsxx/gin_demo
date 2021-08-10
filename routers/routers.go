package routers

import "github.com/gin-gonic/gin"

type Option func(engine *gin.Engine)


var options = []Option{}

// 注册app路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}
//初始化操作
func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options{
		opt(r)
	}
	return r
}
