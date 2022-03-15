//Date:2022/3/15
//FileName:hash
//Author:Silicon-He
//AuthorEmail:silicon_he@163.com
//CreatingTime:10:47

package hash

type HashType int

const (
	BKDRHash HashType = 0
)

type HashFunc func(val interface{}) (int, error)

func NewHashFunc(hashT HashType) HashFunc {
	switch hashT {
	case BKDRHash:
		return bkdr
	default:
		return bkdr
	}
}
