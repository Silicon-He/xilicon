//Date:2022/3/13
//FileName:bitmap
//Author:Silicon-He
//AuthorEmail:silicon_he@163.com
//CreatingTime:20:31

package bitmap

import (
	"errors"
	"fmt"
	"math"
)

type Bitmap struct {
	length int
	bitM   []uint64
}

func NewBitmap() *Bitmap {
	return &Bitmap{
		length: 0,
		bitM:   make([]uint64, 128),
	}
}

func (b *Bitmap) Set(num int) error {
	if num > math.MaxInt32 {
		// https://github.com/golang/go/issues/38673
		// easy get wrong for make a huge slice
		return errors.New(fmt.Sprintf("Not support for length > %d(MaxInt32)", math.MaxInt32))
	}
	offset := num % 64
	pos := num / 64
	diff := pos - len(b.bitM)

	switch {
	case diff > 0 && 2048 > diff:
		b.bitM = append(b.bitM, make([]uint64, diff+64)...)

	case diff >= 2048 && diff <= math.MaxInt32:
		tempBitM := b.bitM
		b.bitM = make([]uint64, pos+64)
		copy(b.bitM, tempBitM)
	}

	if b.bitM[pos]&(1<<offset) == 0 {
		b.bitM[pos] = b.bitM[pos] | (1 << offset)
		b.length++
	}
	return nil
}

func (b *Bitmap) Get(num int) bool {
	offset := num % 64
	pos := num / 64
	return len(b.bitM) > pos && b.bitM[pos]&(1<<offset) != 0
}

func (b *Bitmap) Clean(num int) {
	offset := num % 64
	pos := num / 64
	if pos > len(b.bitM) {
		return
	}
	b.bitM[pos] = b.bitM[pos] & (^uint64(1 << offset))
	b.length--
}

func (b *Bitmap) Length() int {
	return b.length
}
