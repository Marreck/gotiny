package goutils_test

import (
	"testing"
	"time"

	"github.com/Marreck/gotiny/goutils"
)

func TestChannel(t *testing.T) {
	c := goutils.NewChannel(1, time.Second)
	var a = 1
	go c.Add(a)
	select {
	case <-c.FullChan:
		is := c.Get()
		for i := 0; i < len(is); i++ {
			if is[i].(int) != 1 {
				t.Fatalf("error %v", is[i].(int))
			}
		}
	}
}

func TestChannelA(t *testing.T) {
	c := goutils.NewChannel(10, time.Second)
	loopNum := 10000000
	exitch := make(chan struct{})
	go func() {
		for i := 0; i < loopNum; i++ {
			c.Add(i)
		}
		exitch <- struct{}{}
	}()

	n := 0
lb:
	for {
		select {
		case <-c.FullChan:
			r := c.Get()
			for i := 0; i < len(r); i++ {
				if r[i].(int) != n {
					t.Error("error")
				}
				n++
			}
		case <-exitch:
			if c.Len() != 0 {
				go func() { exitch <- struct{}{} }()
				break
			}
			if n != loopNum {
				t.Error("error")
			}
			break lb
		}
	}
}
