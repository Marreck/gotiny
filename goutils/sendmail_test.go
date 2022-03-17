package goutils

import (
	"testing"
)

func TestProcessSendEmail(t *testing.T) {
	err := SendEmail([]string{"402027966@qq.com"}, "test", "test")
	if err != nil {
		t.Error(err)
	}
}
