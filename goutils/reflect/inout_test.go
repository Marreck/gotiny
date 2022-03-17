package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetInAndOut(t *testing.T) {
	fmt.Println(GetInAndOut(reflect.TypeOf(a)))
	in, out := GetInAndOut(reflect.TypeOf(time.Now))
	fmt.Println(reflect.TypeOf((*struct{ Out0 time.Time })(nil)).Elem() == out)
	_ = out
	_ = in
}

func a(i int) int {
	return i
}
