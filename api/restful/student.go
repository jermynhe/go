package restful

import (
	"errors"
	"manger/internal/logic"
	"manger/pkg/misc/resp2"

	"github.com/gin-gonic/gin"
)

// TODO 未实现后面实现

// CreateStudent 添加学生
func (m *Manger) CreateStudent(c *gin.Context) {
	req := new(logic.CreateStudentReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithError(500, err)
		return
	}

	resp, err := m.service.CreateStudent(c, req)
	resp2.Format(c, resp, err)
}

// ListStudentAll 查询所有
func (m *Manger) ListStudentAll(c *gin.Context) {
	resp2.Format(c, nil, errors.New("测试查询所有"))
}

// ListStudent 查询所有
func (m *Manger) ListStudent(c *gin.Context) {
	resp2.Format(c, nil, errors.New("测试查询"))
}

// UpdateStudent 更新学生
func (m *Manger) UpdateStudent(c *gin.Context) {
	resp2.Format(c, nil, errors.New("测试更新"))
}

// DeleteStudent 删除学生
func (m *Manger) DeleteStudent(c *gin.Context) {
	resp2.Format(c, nil, errors.New("测试删除"))
}
