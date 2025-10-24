package utils

// Contain 列表是否包含
func Contain[T comparable](vs []T, v T) bool {
	for i, size := 0, len(vs); i < size; i++ {
		if vs[i] == v {
			return true
		}
	}
	return false
}

// Pcontain 列表是否包含
func Pcontain[T comparable](vs []*T, v *T) bool {
	for i, size := 0, len(vs); i < size; i++ {
		if *(vs[i]) == *v {
			return true
		}
	}
	return false
}

// PmakeX 创建指针元素列表
func PmakeX[T any](size int) []*T {
	rets := make([]*T, size)
	for i := 0; i < size; i++ {
		rets[i] = new(T)
	}
	return rets
}
