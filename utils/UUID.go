package utils

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func NewUUID() string {
	u, err := uuid.NewV4()

	if err != nil {
		fmt.Println("new v4 failed,err:", err)
		return ""
	}

	return u.String()
}
