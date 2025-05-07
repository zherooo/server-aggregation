package utils

import (
	"container/list"
	"errors"
	"fmt"
)

var baseStr string = "ABCD0123456789EFGHJKLMNPQRSTUVWXYZ"
var base []byte = []byte(baseStr)

func Base34Encode(num int) (code string) {
	enNum := uint64(num)
	mod := uint64(0)
	l := list.New()
	for enNum != 0 {
		mod = enNum % 34
		enNum = enNum / 34
		l.PushFront(base[int(mod)])
	}
	listLen := l.Len()
	if listLen >= 6 {
		res := make([]byte, 0, listLen)
		for i := l.Front(); i != nil; i = i.Next() {
			res = append(res, i.Value.(byte))
		}
		return string(res)
	} else {
		res := make([]byte, 0, 6)
		for i := 0; i < 6; i++ {
			if i < 6-listLen {
				res = append(res, base[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}
		}
		return string(res)
	}
}

func Base34Decode(code string) (uint64, error) {
	str := []byte(code)
	baseMap := make(map[byte]int)
	for i, v := range base {
		baseMap[v] = i
	}
	if baseMap == nil {
		return 0, errors.New("no init base map")
	}
	if str == nil || len(str) == 0 {
		return 0, errors.New("parameter is nil or empty")
	}
	var res uint64 = 0
	var r uint64 = 0
	for i := len(str) - 1; i >= 0; i-- {
		v, ok := baseMap[str[i]]
		if !ok {
			fmt.Printf("")
			return 0, errors.New("character is not base")
		}
		var b uint64 = 1
		for j := uint64(0); j < r; j++ {
			b *= 34
		}
		res += b * uint64(v)
		r++
	}
	return res, nil
}
