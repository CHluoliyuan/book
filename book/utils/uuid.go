package utils

import uuid "github.com/satori/go.uuid"

// 设置uuid 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}
