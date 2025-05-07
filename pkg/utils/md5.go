package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}
func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1Encode(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// CalculateFileMD5 自定义函数，用于计算文件的MD5哈希值
func CalculateFileMD5(fileHeader *multipart.FileHeader) (md5Str string, mime string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer file.Close()

	// 创建MD5哈希对象
	hash := md5.New()

	// 读取文件并计算哈希值
	if _, err := io.Copy(hash, file); err != nil {
		return "", "", err
	}

	// 获取MD5哈希值的字节数组
	md5Hash := hash.Sum(nil)

	// 将字节数组格式化为16进制字符串
	md5Str = fmt.Sprintf("%x", md5Hash)

	// 获取MIME
	buffer := make([]byte, 512)
	_, _ = file.Read(buffer)
	mime = http.DetectContentType(buffer)

	return
}

func GetFullExtension(filename string) string {
	dotParts := strings.Split(filename, ".")
	if len(dotParts) > 1 {
		return "." + strings.Join(dotParts[1:], ".")
	}
	return ""
}
