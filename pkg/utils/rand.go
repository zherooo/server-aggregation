package utils

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"
)

const (
	EPOCH       = int64(1514736000000) // 起始时间，设置为开发的时候的时间
	MACHINE_LEN = 5                    // 主机占用 5 个 bits
	PROCESS_LEN = 7                    // 进程 ID 占用 7 个 bits
	RANDOM_LEN  = 10                   // 随机字符串占用 10 个 bits
)

var machineId = 0 // 主机 ID，默认值为 0

// GetRandomString
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成唯一 ID(原来php的uuid生成逻辑搬过来的)
func Gen() (int64, error) {
	// 当前时间的毫秒数
	timeMillis := time.Now().UnixMilli()
	return genByTime(timeMillis)
}

func genByTime(timeMillis int64) (int64, error) {
	// 减掉时间的开始
	timeMillis -= EPOCH

	// 将时间戳转换为二进制字符串
	base := intToBinaryString(timeMillis, 41)
	if len(base) > 41 {
		return 0, fmt.Errorf("时间超出范围")
	}
	base = "0" + base

	var machineIdStr string
	if machineId != 0 {
		machineIdStr = intToBinaryString(int64(machineId), MACHINE_LEN)
		if len(machineIdStr) > MACHINE_LEN {
			return 0, fmt.Errorf("机器 ID 超出范围")
		}
		machineIdStr = padBinaryString(machineIdStr, MACHINE_LEN)
	}

	// 进程 ID 部分
	pid := os.Getpid() % (1 << PROCESS_LEN)
	pidStr := padBinaryString(intToBinaryString(int64(pid), PROCESS_LEN), PROCESS_LEN)

	// 随机字符串部分
	random := rand.Intn(1 << RANDOM_LEN)
	randomStr := padBinaryString(intToBinaryString(int64(random), RANDOM_LEN), RANDOM_LEN)

	// 拼接所有部分
	base = base + machineIdStr + pidStr + randomStr

	// 转换为十进制
	result, success := binaryStringToInt(base)
	if !success {
		return 0, fmt.Errorf("二进制字符串转换失败")
	}

	return result, nil
}

// 将整数转换为二进制字符串，长度为 length
func intToBinaryString(num int64, length int) string {
	binaryStr := ""
	for num > 0 {
		binaryStr = string(num%2+'0') + binaryStr
		num /= 2
	}
	return padBinaryString(binaryStr, length)
}

// 填充二进制字符串到指定长度
func padBinaryString(binaryStr string, length int) string {
	for len(binaryStr) < length {
		binaryStr = "0" + binaryStr
	}
	return binaryStr
}

// 将二进制字符串转换为整数
func binaryStringToInt(binaryStr string) (int64, bool) {
	result, success := new(big.Int).SetString(binaryStr, 2)
	if !success {
		return 0, false
	}
	return result.Int64(), true
}
