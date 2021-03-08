package id2

import uuid "github.com/satori/go.uuid"

// GenerateID id生成器
// 现使用UUID
func GenerateID() string {
	u1 := uuid.Must(uuid.NewV1(), nil)
	return u1.String()
}
