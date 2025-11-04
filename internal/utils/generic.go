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

// Convertm 转化字典
func Convertm[T any, U comparable](vs []T, f Convert[T, U]) map[U]T {
	rets := make(map[U]T, len(vs))
	for _, v := range vs {
		rets[f(v)] = v
	}
	return rets
}

// MapKeys 字典键列表
func MapKeys[T comparable, U any](vs map[T]U) []T {
	rets := make([]T, 0, len(vs))
	for k := range vs {
		rets = append(rets, k)
	}
	return rets
}

// MapValues 字典值列表
func MapValues[T comparable, U any](vs map[T]U) []U {
	rets := make([]U, 0, len(vs))
	for _, v := range vs {
		rets = append(rets, v)
	}
	return rets
}

// Uniques 对切片进行去重
func Uniques[T comparable](args []T) []T {
	if len(args) == 0 {
		return nil
	}
	us := make([]T, 0, len(args))
	for _, s := range args {
		if Contains(us, s) {
			continue
		}
		us = append(us, s)
	}
	return us
}

// Contains 判断切片是否包含元素
func Contains[T comparable](data []T, e T) bool {
	for _, s := range data {
		if s == e {
			return true
		}
	}
	return false
}

// Removes 从切片中移除指定元素
func Removes[T comparable](data []T, es ...T) []T {
	if len(data) == 0 {
		return nil
	}
	us := make([]T, 0, len(data))
	for _, s := range data {
		if Contains(es, s) {
			continue
		}
		us = append(us, s)
	}
	return us
}
