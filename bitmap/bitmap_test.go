//Date:2022/3/14
//FileName:bitmap_test
//Author:Silicon-He
//AuthorEmail:silicon_he@163.com
//CreatingTime:0:02

package bitmap

import (
	"testing"
)

func TestBitmap_Set(t *testing.T) {
	bitmap := NewBitmap()
	bitmap.Set(123)
	t.Log(bitmap.Get(123) == true)
	t.Log(bitmap.Get(159) == false)
	bitmap.Set(159)
	t.Log(bitmap.Get(159) == true)
	t.Log(bitmap.Get(1896557*15632) == false)
	bitmap.Set(1896557*15632)
	t.Log(bitmap.Get(1896557*15632) == true)
	bitmap.Clean(1896557*15632)
	t.Log(bitmap.Get(1896557*15632) == false)
}
