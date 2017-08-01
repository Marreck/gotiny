package gotiny

import (
	"reflect"
	"unsafe"
)

type Decoder struct {
	buf     []byte //buf
	index   int    //下一个要使用的字节在buf中的下标
	boolPos byte   //下一次要读取的bool在buf中的下标,即buf[boolPos]
	boolBit byte   //下一次要读取的bool的buf[boolPos]中的bit位

	decEngs []decEng //解码器集合
	length  int      //解码器数量
}

func Decodes(buf []byte, is ...interface{}) int {
	return NewDecoderWithPtrs(is...).Decodes(buf, is...)
}

func NewDecoderWithPtrs(is ...interface{}) *Decoder {
	l := len(is)
	if l < 1 {
		panic("must have argument!")
	}
	des := make([]decEng, l)
	for i := 0; i < l; i++ {
		rt := reflect.TypeOf(is[i])
		if rt.Kind() != reflect.Ptr {
			panic("must a pointer type!")
		}
		des[i] = getDecEngine(rt.Elem())
	}
	return &Decoder{
		length:  l,
		decEngs: des,
	}
}

func NewDecoder(is ...interface{}) *Decoder {
	l := len(is)
	if l < 1 {
		panic("must have argument!")
	}
	des := make([]decEng, l)
	for i := 0; i < l; i++ {
		des[i] = getDecEngine(reflect.TypeOf(is[i]))
	}
	return &Decoder{
		length:  l,
		decEngs: des,
	}
}

func NewDecoderWithTypes(ts ...reflect.Type) *Decoder {
	l := len(ts)
	if l < 1 {
		panic("must have argument!")
	}
	des := make([]decEng, l)
	for i := 0; i < l; i++ {
		des[i] = getDecEngine(ts[i])
	}
	return &Decoder{
		length:  l,
		decEngs: des,
	}
}

func (d *Decoder) reset() int {
	index := d.index
	d.index = 0
	d.boolPos = 0
	d.boolBit = 0
	return index
}

// is is pointer of variable
func (d *Decoder) Decodes(buf []byte, is ...interface{}) int {
	d.buf = buf
	engs := d.decEngs
	for i := 0; i < d.length; i++ {
		engs[i](d, (*[2]unsafe.Pointer)(unsafe.Pointer(&is[i]))[1])
	}
	return d.reset()
}

// ps is a unsafe.Pointer of the variable
func (d *Decoder) DecodeByUPtr(buf []byte, ps ...unsafe.Pointer) int {
	d.buf = buf
	engs := d.decEngs
	for i := 0; i < d.length; i++ {
		engs[i](d, ps[i])
	}
	return d.reset()
}

func (d *Decoder) DecodeValues(buf []byte, vs ...reflect.Value) int {
	d.buf = buf
	engs := d.decEngs
	for i := 0; i < d.length; i++ {
		engs[i](d, unsafe.Pointer(vs[i].UnsafeAddr()))
	}
	return d.reset()
}
