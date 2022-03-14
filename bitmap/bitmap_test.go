//Date:2022/3/14
//FileName:bitmap_test
//Author:Silicon-He
//AuthorEmail:silicon_he@163.com
//CreatingTime:0:02

package bitmap

import (
	"math"
	"testing"
)

func TestBitmap_Function(t *testing.T) {
	bitmap := NewBitmap()
	bitmap.Set(1)
	t.Log(bitmap.Get(1) == true)
	t.Log(bitmap.Get(math.MaxInt8) == false)
	bitmap.Set(math.MaxInt8)
	bitmap.Set(math.MaxInt16)
	t.Log(bitmap.Get(math.MaxInt16) == true)
	bitmap.Set(math.MaxInt32)
	t.Log(bitmap.Get(math.MaxInt32) == true)
	bitmap.Clean(math.MaxInt32)
	t.Log(bitmap.Get(math.MaxInt32) == false)
	t.Logf(bitmap.Set(math.MaxInt64).Error())
}

func BenchmarkBitmap_Set(b *testing.B) {
	bitmap := NewBitmap()
	for i := 0; i < b.N; i++ {
		bitmap.Set(1)
		bitmap.Set(math.MaxInt8 / 2)
		bitmap.Set(math.MaxInt8)

		bitmap.Set(math.MaxInt16 / 2)
		bitmap.Set(math.MaxInt16)

		bitmap.Set(math.MaxInt32 / 2)
		bitmap.Set(math.MaxInt32)
	}
}

func BenchmarkBitmap_Get(b *testing.B) {
	bitmap := NewBitmap()
	bitmap.Set(1)
	bitmap.Set(math.MaxInt8 / 2)
	bitmap.Set(math.MaxInt8)

	bitmap.Set(math.MaxInt16 / 2)
	bitmap.Set(math.MaxInt16)

	bitmap.Set(math.MaxInt32 / 2)
	bitmap.Set(math.MaxInt32)
	for i := 0; i < b.N; i++ {
		bitmap.Get(1)
		bitmap.Get(math.MaxInt8 / 2)
		bitmap.Get(math.MaxInt8)

		bitmap.Get(math.MaxInt16 / 2)
		bitmap.Get(math.MaxInt16)

		bitmap.Get(math.MaxInt32 / 2)
		bitmap.Get(math.MaxInt32)
	}
}
