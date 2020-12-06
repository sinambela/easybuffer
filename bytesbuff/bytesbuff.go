package bytesbuff

import (
	"bytes"
	"sync"
)

//GetBytesBuffer for
func GetBytesBuffer() (poolx *EasyBytes) {
	poolx = new(EasyBytes)
	(*poolx).init()

	return
}

//EasyBytes for
type EasyBytes struct {
	lock  *sync.Mutex
	poolx *sync.Pool
}

func (x *EasyBytes) init() {
	(*x).lock = new(sync.Mutex)
	(*x).poolx = new(sync.Pool)
	(*x).poolx.New = func() interface{} { return new(bytes.Buffer) }
}

//GetBytesBuffer for
func (x *EasyBytes) GetBytesBuffer() (buff *bytes.Buffer) {
	(*x).lock.Lock()

	buff = (*x).poolx.Get().(*bytes.Buffer)
	buff.Reset()

	(*x).lock.Unlock()

	return
}

//PutBytesBuffer for
func (x *EasyBytes) PutBytesBuffer(buff *bytes.Buffer) {
	(*x).lock.Lock()

	buff.Reset()
	(*x).poolx.Put(buff)

	(*x).lock.Unlock()
}
