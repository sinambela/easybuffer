package keccak256buff

import (
	"hash"
	"sync"

	"golang.org/x/crypto/sha3"
)

//GetKeccak256Buff for
func GetKeccak256Buff() (buffPool *EasyKeccak256) {
	buffPool = new(EasyKeccak256)

	buffPool.init()

	return
}

//EasyKeccak256 for
type EasyKeccak256 struct {
	lock  *sync.Mutex
	poolx *sync.Pool
}

func (x *EasyKeccak256) init() {
	(*x).lock = new(sync.Mutex)

	(*x).poolx = new(sync.Pool)
	(*x).poolx.New = func() interface{} {
		kek256 := sha3.NewLegacyKeccak256()

		return &kek256
	}
}

//GetKeccak256 for getting keccak256
func (x *EasyKeccak256) GetKeccak256() (h *hash.Hash) {
	(*x).lock.Lock()

	h = (*x).poolx.Get().(*hash.Hash)
	(*h).Reset()

	(*x).lock.Unlock()

	return
}

//PutKeccak256 for
func (x *EasyKeccak256) PutKeccak256(h *hash.Hash) {
	(*x).lock.Lock()

	(*h).Reset()
	(*x).poolx.Put(h)

	(*x).lock.Unlock()
}
