package utils

import (
	"reflect"
)

// RemoveDuplicateElement 去除重复的数
func RemoveDuplicateElement(data []int) []int {
	result := make([]int, 0, len(data))
	temp := map[int]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//InArray 是否在数组内
func InArray(item interface{}, array interface{}) bool {
	values := reflect.ValueOf(array)
	if values.Kind() != reflect.Slice {
		return false
	}

	size := values.Len()
	list := make([]interface{}, size)
	slice := values.Slice(0, size)
	for index := 0; index < size; index++ {
		list[index] = slice.Index(index).Interface()
	}

	for index := 0; index < len(list); index++ {
		if list[index] == item {
			return true
		}
	}
	return false
}

func FormatArray(args ...interface{}) (list []interface{}) {
	for _, item := range args {
		slice := reflect.ValueOf(item)
		if slice.Kind() != reflect.Slice {
			list = append(list, item)
			continue
		}

		//slice组合
		items, ok := TakeSliceArg(item)
		if !ok {
			continue
		}
		for _, sliceItem := range items {
			list = append(list, sliceItem)
		}

	}
	return
}
