package reflect

import (
	"reflect"
	"strconv"
)

func GetInAndOut(rt reflect.Type) (in reflect.Type, out reflect.Type) {
	ni, no := rt.NumIn(), rt.NumOut()
	inArg, outArg := make([]reflect.StructField, ni), make([]reflect.StructField, no)
	for i := 0; i < ni; i++ {
		inArg[i].Type = rt.In(i)
		inArg[i].Name = "In" + strconv.Itoa(i)
	}
	in = reflect.StructOf(inArg)

	for i := 0; i < no; i++ {
		outArg[i].Type = rt.Out(i)
		outArg[i].Name = "Out" + strconv.Itoa(i)
	}
	out = reflect.StructOf(outArg)
	return
}

func assignment(a, b interface{}) {

}
