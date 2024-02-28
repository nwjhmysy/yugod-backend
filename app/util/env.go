package util

import (
	"os"
	"strings"
)

// 获取环境变量boolean值
func GetEnvBooleanValue(key string) bool {
	return strings.ToLower(os.Getenv(key)) == "true"
}
