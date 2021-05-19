package main

import (
    "github.com/spaolacci/murmur3"
    //"github.com/willf/bitset"
    "fmt"
    "github.com/bits-and-blooms/bitset"

)

type BloomFilter struct {
    m uint64 // 最大容纳的数
    k uint32 // hash 函数个数
    b *bitset.BitSet



}


func New(m uint64, k uint32) *BloomFilter {
    return &BloomFilter{
        m,k,bitset.New(uint(m)),
    }
}

func (f *BloomFilter) Add(data []byte) {
    for i:=uint32(0);i<f.k;i++ {
        f.b.Set(f.locate(data, i))
    }
}


func (f *BloomFilter) Exist(data []byte) bool {
    for i:=uint32(0);i<f.k;i++ {
        if !f.b.Test(f.locate(data,i)) {
            return false
        }
    }
    return true

}


func (f *BloomFilter) locate(data []byte, seed uint32) uint {
    return getHash(data, seed) % uint(f.m)
}


func getHash(data []byte, seed uint32) uint {
    m := murmur3.New64WithSeed(seed)
    _,_ = m.Write(data)
    return uint(m.Sum64())
}


func main() {
    bloom := New(10000000,3)
    bloom.Add([]byte("a"))
    bloom.Add([]byte("b"))
    bloom.Add([]byte("c"))
    bloom.Add([]byte("asd"))
    fmt.Println(bloom.Exist([]byte("a")))
    fmt.Println(bloom.Exist([]byte("c")))
    fmt.Println(bloom.Exist([]byte("d")))
    fmt.Println(bloom.Exist([]byte("ase")))

}
