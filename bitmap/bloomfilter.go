//Date:2022/3/15
//FileName:bloomfilter
//Author:Silicon-He
//AuthorEmail:silicon_he@163.com
//CreatingTime:0:00

package bitmap

import (
	"xilicon/hash"
)

type BloomFilter struct {
	bitmap    *Bitmap
	bHashFunc []hash.HashFunc
	bHashNum  int
}

// NewBloomFilter for common user
func NewBloomFilter() *BloomFilter {
	return InitBloomFilter(3, hash.BKDRHash)
}

// InitBloomFilter for user who want to customize function
func InitBloomFilter(hashNum int, hashT hash.HashType) *BloomFilter {
	bm := &BloomFilter{
		bitmap:    NewBitmap(),
		bHashFunc: make([]hash.HashFunc, hashNum),
		bHashNum:  hashNum,
	}
	for i := 0; i < bm.bHashNum; i++ {
		bm.bHashFunc[i] = hash.NewHashFunc(hashT)
	}
	return bm
}

func (bf *BloomFilter) CalHash(val interface{}) ([]int, error) {
	bitSlice := make([]int, bf.bHashNum)
	for _, hashFunc := range bf.bHashFunc {
		num, err := hashFunc(val)
		if err != nil {
			return nil, err
		}
		bitSlice = append(bitSlice, num)
	}
	return bitSlice, nil
}

func (bf *BloomFilter) Set(val interface{}) error {
	for _, hashFunc := range bf.bHashFunc {
		num, err := hashFunc(val)
		if err != nil {
			return err
		}
		_, err = bf.bitmap.Set(num)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bf *BloomFilter) Get(val interface{}) (bool, error) {
	for _, hashFunc := range bf.bHashFunc {
		num, err := hashFunc(val)
		if err != nil {
			return false, err
		}
		has := bf.bitmap.Get(num)
		if !has {
			return false, nil
		}
	}
	return true, nil
}
