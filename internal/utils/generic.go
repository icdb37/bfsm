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

// Convert 转换
type Convert[T any, U any] = func(v T) U

// Converts 转化列表
func Converts[T any, U any](vs []T, f Convert[T, U]) []U {
	rets := make([]U, len(vs))
	for i, v := range vs {
		rets[i] = f(v)
	}
	return rets
}
