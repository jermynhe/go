package model

import db "github.com/jinzhu/gorm"

// Student 学生
type Student struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
	Grades string `json:"grades"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	// DeleteSate 0:未删除  1:已删除
	DeleteSate int `json:"delete_sate"`

	DeletedAt int64 `json:"daleted_at"`
}

// StudentRepo 经营标签[存储服务]
type StudentRepo interface {
	// Create 创建标签
	Create(db *db.DB, s *Student) error

	// List 获取标签列表
	List(b *db.DB, id string, name string, page, limit int) ([]*Student, int64, error)

	// Delete 删除标签
	Delete(db *db.DB, id string) error

	// ListAll 获取标签列表
	ListAll(b *db.DB, id string) ([]*Student, error)
}
