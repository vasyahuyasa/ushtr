package hash

import (
	"hash/crc32"
)

type HashFunc func([]byte) string

func NewWithAlphabet(alphabet string, hashlen int) HashFunc {
	enc := []byte(alphabet)
	encLen := uint32(len(enc))
	return func(data []byte) string {
		buf := make([]byte, hashlen)
		h := crc32.ChecksumIEEE(data)
		for i := 0; i < hashlen; i++ {
			buf[i] = enc[(h<<uint((i%4)*8))%encLen]
		}
		return string(buf)
	}
}
