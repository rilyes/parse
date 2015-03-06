package html

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	assert.Equal(t, ToHash([]byte("address")), Address, "'address' must resolve to hash.Address")
	assert.Equal(t, "address", Address.String(), "hash.Address must resolve to 'address'")
	assert.Equal(t, "accept-charset", Accept_Charset.String(), "hash.Accept_Charset must resolve to 'accept-charset'")
}

////////////////////////////////////////////////////////////////

var result int

// naive scenario
func BenchmarkCompareBytes(b *testing.B) {
	var r int
	val := []byte("span")
	for n := 0; n < b.N; n++ {
		if bytes.Equal(val, []byte("span")) {
			r++
		}
	}
	result = r
}

// using-atoms scenario
func BenchmarkFindAndCompareAtom(b *testing.B) {
	var r int
	val := []byte("span")
	for n := 0; n < b.N; n++ {
		if ToHash(val) == Span {
			r++
		}
	}
	result = r
}

// using-atoms worst-case scenario
func BenchmarkFindAtomCompareBytes(b *testing.B) {
	var r int
	val := []byte("zzzz")
	for n := 0; n < b.N; n++ {
		if h := ToHash(val); h == 0 && bytes.Equal(val, []byte("zzzz")) {
			r++
		}
	}
	result = r
}
