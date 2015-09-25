// Package flag provides a simple flag utility
package utils

import (
	"math/big"
)

// const taken from math/big package
const (
	_word0    = big.Word(0)
	_word1    = ^_word0
	_logS     = _word1>>8&1 + _word1>>16&1 + _word1>>32&1
	_S        = 1 << _logS
	_WORDSIZE = _S << 3
)

type Flag struct {
	N    big.Int
	name string
}

var zero big.Int

// Print Flag's name
func (f Flag) String() string {
	return f.name
}

// Check that Flags have the same underlying value
func (f Flag) Cmp(flag Flag) bool {
	return f.N.Cmp(&flag.N) == 0
}

// Initialize a Flag with i-nth bit set to 1
func New(name string, i int) Flag {
	f := Flag{N: big.Int{}, name: name}

	f.N.SetBit(&f.N, i, 1)

	return f
}

// Initialize a Flag that will Intersect with all flags with a value up to (length -1)
func NewBigInt(length int) Flag {
	n := (length + _WORDSIZE - 1) / _WORDSIZE

	wordTable := make([]big.Word, 0)
	for i := 0; i < n; i++ {
		wordTable = append(wordTable, _word1)
	}

	b := big.Int{}
	b.SetBits(wordTable)

	return Flag{N: b}
}

// Counter allows to initialize flags via InitFlag being sure that none of them will Intersect with another one
type Counter uint

// Initialize new Flag
func InitFlag(counter *Counter, name string) Flag {
	f := Flag{N: big.Int{}, name: name}

	f.N.SetBit(&f.N, int(*counter), 1)

	*counter += 1

	return f
}

// Create a new Flag resulting of a Or operation on flags provided
func Join(name string, flags ...Flag) Flag {
	f := Flag{N: big.Int{}, name: name}

	for _, flag := range flags {
		f.N = *(f.N.Or(&f.N, &flag.N))
	}
	return f
}

func Exclude(flag Flag, exclusions ...Flag) Flag {
	for _, exclusion := range exclusions {
		not := new(big.Int).Not(&exclusion.N)
		flag.N = *(flag.N.And(&flag.N, not))
	}

	return flag
}

// Check if both flags intersect
func Intersect(a, b Flag) bool {
	return new(big.Int).And(&a.N, &b.N).Cmp(&zero) != 0
}
