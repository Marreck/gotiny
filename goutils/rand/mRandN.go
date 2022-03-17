// 该程序实现在m个int中选取n个
package rand

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/Marreck/gotiny/goutils"
)

type MrandN struct {
	lock   sync.Mutex
	bucket []int
	bktLen int
	rand   *rand.Rand
}

const defaultLength = 1024

var defaultRand = NewMrandNWithLen(defaultLength)

func NewMrandN(bkt []int) *MrandN {
	return &MrandN{bktLen: len(bkt), bucket: bkt, rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func NewMrandNWithLen(l int) *MrandN {
	bkt := make([]int, l)
	for i := 0; i < l; i++ {
		bkt[i] = i
	}
	return &MrandN{bktLen: l, bucket: bkt, rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (mrn *MrandN) SelectN(n int, data []int) []int {
	if mrn.bktLen < n {
		panic("bucket length " + strconv.Itoa(mrn.bktLen) + " not enough " + strconv.Itoa(n))
	}

	bkt := mrn.bucket
	bl := mrn.bktLen
	r := mrn.rand
	last := mrn.bktLen - 1

	buf, l := goutils.EnlargeInt(data, n)
	temp := buf[l:]

	mrn.lock.Lock()
	for i := 0; i < n; i++ {
		idx := r.Intn(bl - i)
		temp[i] = bkt[idx]
		bkt[idx], bkt[last-i] = bkt[last-i], bkt[idx]
	}
	mrn.lock.Unlock()
	return buf[:l+n]
}

func (mrn *MrandN) Perm(data []int) []int {
	return mrn.SelectN(mrn.bktLen, data)
}

func MSelectN(m, n int, data []int) []int {
	var r = defaultRand
	if m > defaultLength {
		r = NewMrandNWithLen(m)
	}
	return r.SelectN(n, data)
}

func Perm(n int, data []int) []int {
	return MSelectN(n, n, data)
}
