package storage

import (
	"sync/atomic"
	"bytes"
	"errors"
)

var slugAlphabet = [...]byte {
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
        'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
        '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
};

// helper function to calculate pow for ints
func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func idToSlug(id int) (slug string, err error) {
	if id <= 0 {
		return "", errors.New("id must be positive integer")
	}

	slugBytes := []byte{}

	remainder := 0
	dividend := id

	for dividend > 0 {
		remainder = dividend % 62
		dividend = dividend / 62

		slugBytes = append([]byte{ slugAlphabet[remainder] }, slugBytes...)
	}

	return string(slugBytes[:len(slugBytes)]), nil
}

func slugToId(slug string) (id int, err error) {
	tmpId := 0

	slugBytes := []byte(slug)
	slugBytesLen := len(slugBytes)

	for i, r := range slugBytes {
		exponent := slugBytesLen - 1 - i
		multiplicator := bytes.IndexByte(slugAlphabet[:], r)

		if (multiplicator == -1) {
			return 0, errors.New("Invalid character found in slug")
		}

		tmpId += multiplicator * pow(62, exponent)
	}

	return tmpId, nil
}



// @TODO rename count32 to UrlId or so
var idCounter int32

func StoreUrl(url string) {

}

func LoadUrl(slug string) {

}
