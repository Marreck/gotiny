package rand

import (
	"math/rand"
	"time"

	"github.com/Marreck/gotiny/goutils"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	NUM = 1 << iota
	LOWER
	UPPER

	ALL = NUM | LOWER | UPPER

	num   = "1234567890"
	lower = "abcdefghijklmnopqrstuvwxyz"
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	chars = [...][]byte{
		0:                   []byte{},                    // 0
		NUM:                 []byte(num),                 // 001
		LOWER:               []byte(lower),               // 010
		UPPER:               []byte(upper),               // 100
		LOWER | NUM:         []byte(num + lower),         // 011
		UPPER | NUM:         []byte(num + upper),         // 101
		LOWER | UPPER:       []byte(lower + upper),       // 110
		NUM | LOWER | UPPER: []byte(num + upper + lower), // 111
	}
	lens = [...]int{
		0:                   0,
		NUM:                 len(num),
		LOWER:               len(lower),
		UPPER:               len(upper),
		LOWER | NUM:         len(num + lower),
		UPPER | NUM:         len(num + upper),
		LOWER | UPPER:       len(lower + upper),
		NUM | LOWER | UPPER: len(num + upper + lower),
	}
)

//
func BytesToBuff(l int, buff []byte) []byte {
	for i := 0; i < l; i++ {
		buff[i] = byte(r.Intn(1 << 8))
	}
	return buff
}

func Bytes(l int) []byte {
	return BytesToBuff(l, nil)
}

func RandStringAWithChars(l int, chars []byte, data []byte) []byte {
	buf, n := goutils.EnlargeByte(data, l)
	ret := buf[n:]
	for i := 0; i < l; i++ {
		ret[i] = chars[r.Intn(len(chars))]
	}
	return buf
}

func RandStringWithChars(l int, chars []byte) []byte {
	return RandStringAWithChars(l, chars, nil)
}

func RandStringAWithType(l int, typ int, data []byte) []byte {
	return RandStringAWithChars(l, chars[typ], data)
}
func RandStringWithType(l int, typ int) []byte {
	return RandStringAWithType(l, typ, nil)
}

func RandString(l int) []byte {
	return RandStringWithType(l, ALL)
}
func RandStringA(l int, data []byte) []byte {
	return RandStringAWithType(l, ALL, data)
}
