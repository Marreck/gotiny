package rand

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Marreck/gotiny/goutils"
)

func TestMRandN(t *testing.T) {
	for i := 0; i < 10; i++ {
		rst := Perm(i, nil)
		if goutils.CheckRepeat(rst) {
			t.Error(rst)
		}
	}
}

func BenchmarkMSelectN(b *testing.B) {
	r := NewMrandNWithLen(100000)
	for i := 0; i < b.N; i++ {
		r.SelectN(i, nil)
	}
}

func BenchmarkPerm(b *testing.B) {
	buf := make([]int, 0, 100000)
	r := NewMrandNWithLen(100000)
	for i := 0; i < b.N; i++ {
		r.Perm(buf)
	}
}

func BenchmarkMathRand(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		r.Perm(100000)
	}
}
