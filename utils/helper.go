package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

func GenUUID(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, uuid.NewV4().String())
}

func StringToInt64(str string) int64 {
	tmp, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return tmp
}
