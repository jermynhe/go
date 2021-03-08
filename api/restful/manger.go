package restful

import (
	"manger/internal/logic"
	"manger/pkg/misc/config"

	"github.com/gin-gonic/gin"
)

// Manger 管理
type Manger struct {
	group   *gin.RouterGroup
	conf    *config.Config
	service logic.MangerService
}

// NewManger 开启标书路由
func NewManger(conf *config.Config, group *gin.RouterGroup) (*Manger, error) {
	service, err := logic.New(conf)
	if err != nil {
		return nil, err
	}
	return &Manger{
		group:   group,
		conf:    conf,
		service: service,
	}, nil

}

func (m *Manger) router() error {

	// 通用端接口
	s := m.group.Group("/api/v1/student")
	{
		// 获取全部学生
		s.GET("/select/all", m.ListStudentAll)

		// 获取指定id的学生
		s.GET("/select/single/:id", m.ListStudent)

		// 创建学生
		s.POST("/create", m.CreateStudent)

		// 修改学生
		s.PUT("/update", m.UpdateStudent)

		// 删除学生
		s.DELETE("/delete/:id", m.DeleteStudent)
	}
	return nil
}

// Close 关闭标书服务
func (m *Manger) Close() error {
	return m.service.Close()
}
