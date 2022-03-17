package goutils

type Week52 struct {
	store  []float64
	high   uint32
	low    uint32
	start  uint32
	mode   uint32
	length uint32
}

func NewWeek52(cap int) *Week52 {
	length := minQuantity(uint32(cap))
	return &Week52{
		store:  make([]float64, length),
		high:   length - 1,
		low:    0,
		start:  0,
		mode:   length - 1,
		length: length,
	}

}

// round 到最近的2的倍数
func minQuantity(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
