package goutils

import (
	"unicode/utf8"
)

func encodeutf8(p []byte) []byte {
	var buf = make([]byte, 0, 100)
	for len(p) > 0 {
		r, l := utf8.DecodeRune(p)
		if r != utf8.RuneError {
			buf = encUint32(uint32(r), buf)
		}
		p = p[l:]
	}
	return buf
}

func decodeutf8(buf []byte) []byte {
	var p = make([]byte, 10000)
	sl := 0
	for len(buf) > 0 {
		r, l := decUint32(buf)
		if r != utf8.RuneError {
			sl += utf8.EncodeRune(p[sl:], rune(r))
		}
		buf = buf[l:]
	}
	return p[:sl]
}

func encUint32(v uint32, buf []byte) []byte {
	switch {
	case v < 1<<7-1:
		buf = append(buf, byte(v))
	case v < 1<<14-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7))
	case v < 1<<21-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14))
	case v < 1<<28-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21))
	default:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28))
	}
	return buf
}

func decUint32(buf []byte) (uint32, int) {
	i := 0
	x := uint32(buf[i])
	if x < 0x80 {
		i++
		return x, i
	}
	x1 := buf[i+1]
	x += uint32(x1) << 7
	if x1 < 0x80 {
		return x - 1<<7, i + 2
	}
	x2 := buf[i+2]
	x += uint32(x2) << 14
	if x2 < 0x80 {
		return x - (1<<7 + 1<<14), i + 3
	}
	x3 := buf[i+3]
	x += uint32(x3) << 21
	if x3 < 0x80 {
		return x - (1<<7 + 1<<14 + 1<<21), i + 4
	}
	x4 := buf[i+4]
	x += uint32(x4) << 28
	return x - (1<<7 + 1<<14 + 1<<21 + 1<<28), i + 5
}
