package postgres

import (
	"manger/internal/model"
	"strings"

	"github.com/jinzhu/gorm"
	db "github.com/jinzhu/gorm"
)

// NewStudentRepo 创建一个标签存储服务
func NewStudentRepo() model.StudentRepo {
	return &studentRepo{}
}

type studentRepo struct {
}

func (t *studentRepo) TableName() string {
	return "tbl_manger_student"
}

// Create 创建标签
func (t *studentRepo) Create(db *db.DB, Student *model.Student) error {
	return db.Table(t.TableName()).Create(Student).Error
}

// List 获取标签列表
func (t *studentRepo) List(db *db.DB, ID string, Student string, page, limit int) ([]*model.Student, int64, error) {
	ql := db.Table(t.TableName())

	ql = ql.Where("id = ?", ID)
	if Student != "" {
		Student = strings.Trim(Student, "%")
		// 后模糊
		ql = ql.Where("Student like '" + Student + "%'")
	}

	var total int64
	ql.Count(&total)

	if limit != 0 {
		ql = ql.Offset((page - 1) * limit).Limit(limit)
	}
	ql = ql.Order("created_at DESC")

	Students := make([]*model.Student, 0)

	err := ql.Find(&Students).Error
	return Students, total, err
}

// Delete 删除标签
func (t *studentRepo) Delete(db *db.DB, id string) error {
	ql := db.Table(t.TableName()).Where("id = ?", id)

	return ql.Delete(model.Student{}).Error
}

// ListAll 获取标签列表
func (t *studentRepo) ListAll(db *db.DB, ID string) ([]*model.Student, error) {
	ql := db.Table(t.TableName())

	ql = ql.Where("id = ?", ID)

	ql = ql.Order("created_at DESC")

	Students := make([]*model.Student, 0)

	err := ql.Find(&Students).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return Students, err
}

// Close 关闭服务
func (t *studentRepo) Close() error {
	return nil
}
