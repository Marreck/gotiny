package net

import (
	"fmt"
	"testing"
)

func TestGetLocalIPAddrs(t *testing.T) {
	ips, err := GetLocalIPAddrs()
	if err == nil {
		for i := 0; i < len(ips); i++ {
			fmt.Println(ips[i].String())
		}
	}
}
