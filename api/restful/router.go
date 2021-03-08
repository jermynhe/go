package restful

import (
	"manger/pkg/misc/config"

	logger "manger/pkg/misc/log"

	"github.com/gin-gonic/gin"
)

// Router 路由
type Router struct {
	c      *config.Config
	e      *gin.Engine
	manger *Manger
}

// NewRouter 创建路由集
func NewRouter(c *config.Config) (*Router, error) {
	e := gin.New()

	r := &Router{
		c: c,
		e: e,
	}

	r.mid()

	group := e.Group("/manger")

	manger, err := NewManger(c, group)
	if err != nil {
		return nil, err
	}
	r.manger = manger

	if err := r.router(); err != nil {
		return nil, err
	}

	return r, nil
}

// 中间件
func (r *Router) mid() {
	r.e.Use(gin.LoggerWithConfig(
		gin.LoggerConfig{
			Output: logger.Logger.Out,
		}), gin.Recovery())
}

func (r *Router) router() error {

	return r.manger.router()
}

// Run 开启restful服务
func (r *Router) Run() error {
	return r.e.Run(r.c.HTTP.Addr)
}

// Close 关闭路由
func (r *Router) Close() error {
	return r.manger.Close()
}
