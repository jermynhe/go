package model

import "time"

const (
	// CounterGoods 物品计数器
	CounterGoods = "goods"

	// CounterPaper 标书计数器
	CounterPaper = "paper"
)

const (
	// CounterDayTTL day 计数器过期时间
	CounterDayTTL = time.Second * 60 * 60 * 24
)

const (
	// CounterDayFormat day格式
	CounterDayFormat = "20060102"
)

// Counter 计数器
type Counter struct {
	Name  string
	Count int64
	TTL   time.Duration
}

// CounterRepo 计数器[存储服务]
type CounterRepo interface {

	// Count 计数
	Count(*Counter) (int64, error)
}
