package xorfilter

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	testsize := 10000
	keys := make([]uint64, testsize, testsize)
	for i := range keys {
		keys[i] = rand.Uint64()
	}
	filter := Populate(keys)
	for _, v := range keys {
		assert.Equal(t, true, filter.Contains(v))
	}
	falsesize := 1000000
	matches := 0
	bpv := float64(len(filter.Fingerprints)) * 8.0 / float64(testsize)
	fmt.Println("bits per entry ", bpv)
	assert.Equal(t, true, bpv < 10.)
	for i := 0; i < falsesize; i++ {
		v := rand.Uint64()
		if filter.Contains(v) {
			matches++
		}
	}
	fpp := float64(matches) * 100.0 / float64(falsesize)
	fmt.Println("false positive rate ", fpp)
	assert.Equal(t, true, fpp < 0.40)

}
