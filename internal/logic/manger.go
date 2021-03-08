package logic

import (
	"context"

	"manger/internal/model"
	"manger/internal/model/postgres"
	redis2 "manger/internal/model/redis"
	"manger/pkg/misc/config"
	logger "manger/pkg/misc/log"

	"github.com/go-redis/redis"
	db "github.com/jinzhu/gorm"
)

// MangerService 标书接口
type MangerService interface {
	// CreateStudent 创建标签
	CreateStudent(c context.Context, req *CreateStudentReq) (resp *CreateStudentResp, err error)

	// // ListStudent 获取标签列表
	// ListStudent(ctx context.Context, req *ListStudentReq) (resp *ListStudentResp, err error)

	// // ListStudentAll 获取标签列表
	// ListStudentAll(ctx context.Context, req *ListStudentAllReq) (resp *ListStudentAllResp, err error)

	// // DeleteStudent 删除标签
	// DeleteStudent(ctx context.Context, req *DeleteStudentReq) (resp *DeleteStudentResp, err error)

	// Close 关闭服务
	Close() error
}

// New 创建管理逻辑服务
func New(conf *config.Config) (MangerService, error) {
	db, err := postgres.New(conf.DB.Postgres)
	if err != nil {
		return nil, err
	}

	student := postgres.NewStudentRepo()

	redis, err := redis2.New(&conf.DB.Redis)

	manger := &manger{
		conf:    conf,
		db:      db,
		redis:   redis,
		student: student,
	}
	return manger, nil
}

type manger struct {
	conf  *config.Config
	db    *db.DB
	redis *redis.Client

	student model.StudentRepo
}

// Close 关闭服务
func (m *manger) Close() error {

	if err := m.redis.Close(); err != nil {
		logger.Logger.Errorf("redis close err:%v", err)
	}
	if err := m.db.Close(); err != nil {
		logger.Logger.Errorf("db close err:%v", err)
	}
	return nil
}
