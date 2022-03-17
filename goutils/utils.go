package goutils

func CheckRepeat(a []int) bool {
	set := NewSet()
	for _, i := range a {
		if set.Has(int64(i)) {
			return true
		} else {
			set.Add(int64(i))
		}
	}
	return false
}

func EnlargeInt(data []int, l int) (buf []int, n int) {
	if cap(data)-len(data) < l {
		buf = make([]int, l+len(data))
		n = copy(buf, data)
	} else {
		buf = data[:len(data)+l]
	}
	return buf, len(data)
}
func EnlargeByte(data []byte, l int) (buf []byte, n int) {
	if cap(data)-len(data) < l {
		buf = make([]byte, l+len(data))
		n = copy(buf, data)
	} else {
		buf = data[:len(data)+l]
	}
	return buf, len(data)
}
