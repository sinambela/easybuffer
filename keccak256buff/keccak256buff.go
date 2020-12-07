package keccak256buff

import (
	"hash"
	"sync"

	"golang.org/x/crypto/sha3"
)

//EasyKeccak256 for
type EasyKeccak256 struct {
	lock  *sync.Mutex
	poolx *sync.Pool
}

func (x *EasyKeccak256) init() {
	(*x).lock = new(sync.Mutex)

	(*x).poolx = new(sync.Pool)
	(*x).poolx.New = func() interface{} {
		kek256 := sha3.NewLegacyKeccak512()

		return &kek256
	}
}

//GetKeccak256 for getting keccak256
func (x *EasyKeccak256) GetKeccak256() (h *hash.Hash) {
	h = (*x).poolx.Get().(*hash.Hash)
	(*h).Reset()

	return
}

//PutKeccak256 for
func (x *EasyKeccak256) PutKeccak256(h *hash.Hash) {
	(*h).Reset()
	(*x).poolx.Put(h)
}
