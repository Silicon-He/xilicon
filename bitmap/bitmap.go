//Date:2022/3/13
//FileName:bitmap
//Author:Silicon-He
//AuthorEmail:silicon_he@163.com
//CreatingTime:20:31

package bitmap

type Bitmap struct {
	length int32
	bitM []uint64
}

func NewBitmap() *Bitmap {
	return &Bitmap{
		length: 0,
		bitM: []uint64{},
	}
}

func (b *Bitmap) Set(num int) {
	offset := num % 64
	pos := num / 64
	for pos >= len(b.bitM){
		b.bitM = append(b.bitM,0)
	}
	if b.bitM[pos] & (1 << offset) == 0 {
		b.bitM[pos] = b.bitM[pos] | (1 << offset)
		b.length ++
		return
	}
}

func (b *Bitmap) Get(num int) bool {
	offset := num % 64
	pos := num / 64
	return len(b.bitM) > pos && b.bitM[pos] & (1 << offset) != 0
}

func (b *Bitmap)Clean(num int)  {
	offset := num % 64
	pos := num / 64
	if pos > len(b.bitM) {
		return
	}
	b.bitM[pos] = b.bitM[pos] & (^uint64(1 << offset))
	return
}