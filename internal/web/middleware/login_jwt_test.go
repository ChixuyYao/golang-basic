package middleware

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	url := "/users/signup"
	allowURL := []string{"/users/signup", "/users/login"}
	for _, v := range allowURL {
		if url == v {
			return
		}
	}
	fmt.Println("=======Pass=======")
}
