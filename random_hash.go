package randomhash

import (
	"bytes"
	"crypto/rand"
	"math"
	"strconv"
	"unicode/utf8"
)

//RandomHash RandomHash
type RandomHash struct {
	Charset    []rune
	charLength int
}

//charLength calc the bit count for rune char
func charLength(charset string) int {

	charsetLength := utf8.RuneCountInString(charset)
	if charsetLength == 0 {
		panic("invalid charset")
	}

	if (charsetLength & (charsetLength - 1)) != 0 {
		panic("charset needs to contain exactly 2^n characters.")
	}

	bstr := strconv.FormatInt(int64(charsetLength), 2)

	return len(bstr) - 1

}

// New get new RandomHash
func New(charset string) *RandomHash {

	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	}
	charLength := charLength(charset)
	rh := &RandomHash{Charset: []rune(charset), charLength: charLength}

	return rh
}

// GenerateHash generate the length hash string
func (randomHash *RandomHash) GenerateHash(length int) (string, error) {
	numBytes := math.Ceil(float64(length*randomHash.charLength) / float64(8))

	b := make([]byte, int(numBytes))
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return randomHash.encoding(b[:]), nil
}

//encoding encoding the bytes data
func (randomHash *RandomHash) encoding(data []byte) string {
	numBits := len(data) * 8
	pos := 0

	var res bytes.Buffer
	for pos+randomHash.charLength <= numBits {

		chunk := 0
		for i := 0; i < randomHash.charLength; i++ {
			_pos := pos + i
			b := data[_pos>>3] // pos/8
			_pos = (_pos & 7)  // pos%8
			bit := (int(b) >> uint(7-_pos)) & 1

			chunk <<= 1
			chunk |= bit
		}

		res.WriteRune(randomHash.Charset[chunk])
		pos += randomHash.charLength
	}

	return res.String()
}
