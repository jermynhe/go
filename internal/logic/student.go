package logic

import (
	"context"
	"manger/internal/model"
	"manger/pkg/misc/id2"
	"manger/pkg/misc/time2"
)

// CreateStudentReq 创建学生参数
type CreateStudentReq struct {
	Name       string `json:"name"`
	Age        string `json:"age"`
	Grades     string `json:"grades"`
	DeleteSate int    `json:"delete_sate"`
}

// CreateStudentResp 创建标签返回值
type CreateStudentResp struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Age        string `json:"age"`
	Grades     string `json:"grades"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	DeleteSate int    `json:"delete_sate"`
	DeletedAt  int64  `json:"daleted_at"`
}

// CreateStudent 创建学生
func (m *manger) CreateStudent(c context.Context, req *CreateStudentReq) (resp *CreateStudentResp, err error) {
	s := &model.Student{
		ID:         id2.GenerateID(),
		Name:       req.Name,
		Grades:     req.Grades,
		Age:        req.Age,
		CreatedAt:  time2.NowUTCUnix(),
		UpdatedAt:  time2.NowUTCUnix(),
		DeleteSate: req.DeleteSate,
		DeletedAt:  0,
	}
	tx := m.db.Begin()
	if err := m.student.Create(tx, s); err != nil {
		tx.Rollback()
		return resp, err
	}
	tx.Commit()
	// 返回值
	resp = &CreateStudentResp{
		ID:         s.ID,
		Name:       s.Name,
		Grades:     s.Grades,
		Age:        s.Age,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
		DeleteSate: s.DeleteSate,
		DeletedAt:  s.DeletedAt,
	}
	return
}
