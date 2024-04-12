package util

type IntInterface interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

// 切片去重
func Unique[T IntInterface | string](data []T) []T {

	list, hash := make([]T, 0), make(map[T]struct{})

	for _, value := range data {
		if _, ok := hash[value]; !ok {
			list = append(list, value)
			hash[value] = struct{}{}
		}
	}

	return list
}

// 切片转 Map
func ToMap[T any, K int | string](arr []T, fn func(T) K) map[K]T {

	var m = make(map[K]T)
	for _, t := range arr {
		m[fn(t)] = t
	}

	return m
}

func Delete(slice []string, index int) (newSlice []string) {

	if index < 0 || index >= len(slice) {
		return slice
	}

	if index == 0 {
		return slice[1:]
	} else if index == len(slice)-1 {
		return slice[:index]
	}

	return append(slice[:index], slice[index+1:]...)
}
