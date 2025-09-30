package utils

import (
	"fmt"
	"testing"
)

type Stack[T any] struct{}

func TestFetchx(t *testing.T) {
	k1 := Stack[int]{}
	k2 := Stack[int]{}
	ms := map[any]int{}
	ms[k1] = 1
	ms[k2] = 2
	s1 := Stack[string]{}
	s2 := Stack[string]{}
	ms[s1] = 3
	ms[s2] = 4
	fmt.Println(k1, k2, ms)
}
