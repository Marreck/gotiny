package goutils

import (
	"sync"
	"time"
)

type Channel struct {
	exitCh    chan struct{}
	mx        sync.Mutex
	cache     []interface{}
	cache2    []interface{}
	fullNum   int
	FullChan  chan struct{}
	isNoticed bool
}

func NewChannel(fullNum int, interval time.Duration) *Channel {
	c := &Channel{
		fullNum:  fullNum,
		cache:    make([]interface{}, 0, fullNum*2),
		cache2:   make([]interface{}, 0, fullNum*2),
		FullChan: make(chan struct{}),
		exitCh:   make(chan struct{}),
	}
	exitCh := c.exitCh
	go func() {
		for {
			select {
			case <-exitCh:
				return
			default:
				time.Sleep(interval)
				if len(c.cache) > 0 {
					c.FullChan <- struct{}{}
				}
			}
		}
	}()
	return c
}

func (c *Channel) Add(msg ...interface{}) {
	var needNotice bool
	c.mx.Lock()
	c.cache = append(c.cache, msg...)
	needNotice = len(c.cache) >= c.fullNum && !c.isNoticed
	if needNotice {
		c.isNoticed = true
	}
	c.mx.Unlock()
	if needNotice {
		c.FullChan <- struct{}{}
	}
}

func (c *Channel) Get() (ret []interface{}) {
	c.mx.Lock()
	c.cache, c.cache2 = c.cache2, c.cache
	c.cache = c.cache[:0]
	c.isNoticed = false
	c.mx.Unlock()
	return c.cache2
}

func (c *Channel) Len() int {
	c.mx.Lock()
	l := len(c.cache)
	c.mx.Unlock()
	return l
}

func (c *Channel) Close() error {
	close(c.exitCh)
	return nil
}
