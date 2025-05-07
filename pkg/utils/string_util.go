package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	Valid_FirstLicenceTime = "^((19[0-9]{2}|20[0-3][0-9])-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|((([1-2][0-9])(0[48]|[2468][048]|[13579][26]))-02-29)$"
	Valid_Name             = "^[\u4e00-\u9fa5]+$"
)

//Left 字串截取
func Left(s string, length int) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	if len(runes) > length {
		return string(runes[0:length]) + ".."
	}
	return s
}
func GetFileSuffix(s string) string {
	re, _ := regexp.Compile(".(jpg|jpeg|png|gif|exe|doc|docx|ppt|pptx|xls|xlsx)")
	suffix := re.ReplaceAllString(s, "")
	return suffix
}

func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		return RandInt64(min, max)
	}
	return i.Int64()
}

// RandInt [min, max]
func RandInt(min, max int) int {
	return int(RandInt64(int64(min), int64(max+1)))
}

func Strim(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}

func Unicode(rs string) string {
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16)
		}
	}
	return json
}

func HTMLEncode(rs string) string {
	html := ""
	for _, r := range rs {
		html += "&#" + strconv.Itoa(int(r)) + ";"
	}
	return html
}

//Utf8ToGbk utf8编码转gbk
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//GbkToUtf8 gbk编码转utf8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

const (
	regular = "^1\\d{10}$"
)

func ValidatePhone(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func HideMobile(mobileNum string) (HideMobileNum string) {
	if len(mobileNum) < 11 {
		return mobileNum
	}
	HideMobileNum = Substr(mobileNum, 0, 3) + fmt.Sprint("*****", Substr(mobileNum, len(mobileNum)-3, 3))
	return
}

func RegMatch(str string, reg string) bool {
	r := regexp.MustCompile(reg)
	return r.MatchString(str)
}
func RegFindString(str string, key string, reg string) (result string) {
	r := regexp.MustCompile(reg)
	m := r.FindStringSubmatch(str)
	if m != nil {
		for i, name := range r.SubexpNames() {
			if name == key {
				return m[i]
			}
		}
	}
	return
}

// ValidateBankCardID 银行卡号验证
func ValidateBankCardID(cardId string) bool {
	var oddSum int
	var evenSum int

	oddSum = 0
	evenSum = 0

	cardId = reverseString(cardId)

	for i, _ := range cardId {

		item, _ := strconv.Atoi(string(cardId[i]))
		//fmt.Println("", item)

		if i%2 == 1 {
			num := item * 2
			sum := num
			if num > 9 {
				sum = 0
				for _, n := range strconv.Itoa(num) {
					inum, _ := strconv.Atoi(string(n))
					sum += inum
				}
			}
			evenSum += sum
		} else {
			oddSum += item
		}
		//fmt.Println(oddSum, "-----", evenSum)
		//fmt.Println("--------------")
	}

	//fmt.Println("", oddSum, evenSum)
	return (oddSum+evenSum)%10 == 0
}

// reverseString字符串反转
func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// FromIntArray 将[]int数组转为","分隔的字符串
// 例如 [1,2,3] 转换为 "1,2,3"
func FromIntArray(arr []int) string {
	n := len(arr)
	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		strArr[i] = strconv.Itoa(arr[i])
	}
	return strings.Join(strArr, ",")
}

// FromStringArray 将数组转为","分隔的字符串
// 例如 [1,2,3] 转换为 "1,2,3"
func FromStringArray(arr []string) string {
	n := len(arr)
	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		strArr[i] = arr[i]
	}
	return strings.Join(strArr, ",")
}

// FromInt64Array 将[]int数组转为","分隔的字符串
// 例如 [1,2,3] 转换为 "1,2,3"
func FromInt64Array(arr []int64) string {
	n := len(arr)
	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		strArr[i] = strconv.FormatInt(arr[i], 10)
	}
	return strings.Join(strArr, ",")
}

// GetIntArray 将","分隔的字符串转换为数字数组
// 例如 "1,2,3" 转换为 [1, 2, 3]
func GetIntArray(str string) (ret []int, err error) {
	for _, item := range strings.Split(str, ",") {
		if item != "" {
			var value int
			value, err = strconv.Atoi(item)
			if err != nil {
				return
			}
			ret = append(ret, value)
		}
	}
	return
}

// GetStringArray 将","分隔的字符串转换为字符串数组
// 例如 "1,2,3" 转换为 ["1", "2", "3"]
func GetStringArray(str string) (ret []string) {
	for _, item := range strings.Split(str, ",") {
		if item != "" {
			ret = append(ret, item)
		}
	}
	return
}

// GetfloatArray 将","分隔的字符串转换为数字数组
// 例如 "1.1,2,3" 转换为 [1.1, 2, 3]
func Getfloat64Array(str string) (ret []float64, err error) {
	for _, item := range strings.Split(str, ",") {
		if item != "" {
			var value float64
			value, err = strconv.ParseFloat(item, 64)
			if err != nil {
				return
			}
			ret = append(ret, value)
		}
	}
	return
}

// GetInt64Array 将","分隔的字符串转换为数字数组
// 例如 "1,2,3" 转换为 [1, 2, 3]
func GetInt64Array(str string) (ret []int64, err error) {
	for _, item := range strings.Split(str, ",") {
		if item != "" {
			var value int64
			value, err = strconv.ParseInt(item, 10, 64)
			if err != nil {
				return
			}
			ret = append(ret, value)
		}
	}
	return
}

//ReplaceSpecialCharacters 过滤特殊字符
func ReplaceSpecialCharacters(str string) string {
	if len(str) > 0 {
		reg := regexp.MustCompile("[';%]")
		str = reg.ReplaceAllString(str, "")
	}
	return str
}

//RegReplace 正则替换
func RegReplace(str string, reg string) string {
	if len(str) > 0 {
		reg := regexp.MustCompile(reg)
		str = reg.ReplaceAllString(str, "")
	}
	return str
}

//ReplaceSpecialCharacters 过滤特殊字符
func CheckIsIncludeNumOrLetter(str string) bool {
	ok, _ := regexp.MatchString("^[0-9a-zA-Z_]{1,}$", str)
	return ok
}

// dial using TLS/SSL
func dial(addr string) (*tls.Conn, error) {
	return tls.Dial("tcp", addr, nil)
}

// compose message according to "from, to, subject, body"
func composeMsg(from string, to string, subject string, body string) (message string) {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject
	headers["Content-type"] = "text/html;charset=utf-8"
	// Setup message
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	return
}

//对比相同struct下不同的值，修改日志专用，返回map类型
func Contrast(DataOld, DataNew interface{}) string {

	m := ""
	old := reflect.TypeOf(DataOld)
	new := reflect.TypeOf(DataNew)
	if old != new {
		return m
	}
	oldVal := reflect.ValueOf(DataOld)
	newVal := reflect.ValueOf(DataNew)
	nums := old.NumField()
	for i := 0; i < nums; i++ {
		if old.Field(i).Name == new.Field(i).Name && oldVal.Field(i).Interface() != newVal.Field(i).Interface() {

			switch reflect.TypeOf(oldVal.Field(i).Interface()).String() {
			//多选语句switch
			case "string":
				if newVal.Field(i).Interface() != nil && newVal.Field(i).Interface() != "" {
					m = m + old.Field(i).Name + ":" + fmt.Sprintf("%s -> %s", oldVal.Field(i).Interface(), newVal.Field(i).Interface()) + ";<br/>"
				}

			case "int":
				if newVal.Field(i).Interface() != nil {
					m = m + old.Field(i).Name + ":" + fmt.Sprintf("%d -> %d", oldVal.Field(i).Interface(), newVal.Field(i).Interface()) + ";<br/>"
				}
			case "float64":
				if newVal.Field(i).Interface() != nil {
					m = m + old.Field(i).Name + ":" + fmt.Sprintf("%f -> %f", oldVal.Field(i).Interface(), newVal.Field(i).Interface()) + ";<br/>"
				}
			}

		}
	}
	return m

}
func CheckPassword(pwd string) bool {
	if len(pwd) < 8 || len(pwd) > 20 {
		return false
	}
	intHave := false
	lowerHave := false
	PowerHave := false
	for index := 0; index < len(pwd); index++ {
		c := pwd[index]
		if c >= 'a' && c <= 'z' { // 包含a-z
			lowerHave = true
		} else if c >= 'A' && c <= 'Z' { // 包含A-Z
			PowerHave = true
		} else if c >= '0' && c <= '9' { // 包含0-9
			intHave = true
		}
	}
	if (!lowerHave && !PowerHave) || (!lowerHave && !intHave) || (!PowerHave && !intHave) {
		return false
	}
	return true
}

func ValidCarNumber(CarNumber string) bool {
	carNumberReg := "^([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z]{1}[A-Z]{1}(([0-9]{5}[DF])|([DF]([A-HJ-NP-Z0-9])[0-9]{4})))|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z]{1}[A-Z]{1}[A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳]{1})$"
	reg := regexp.MustCompile(carNumberReg)
	return reg.MatchString(CarNumber)
}

// ValidRemarks 校验是否为空
func ValidRemarks(remarks string) bool {
	if strings.TrimSpace(remarks) == "" {
		return false
	}
	return true
}

// ValidPhone 校验是否是合法的手机号 正常手机号 + 110开头
func ValidPhone(phone string) bool {
	if phone == "" {
		return true
	}
	reg := `^1(1|3|4|5|6|7|8|9)\d{9}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

// 正则验证
func ValidString(validReg, str string) bool {
	reg := regexp.MustCompile(validReg)
	return reg.MatchString(str)
}
func TakeSliceArg(arg interface{}) (out []interface{}, ok bool) {

	slice := reflect.ValueOf(arg)
	if slice.Kind() != reflect.Slice {
		return nil, false
	}

	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

//去除双引号
func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// TransString2Map
func TransString2Map(args string) (data map[int]string) {
	data = make(map[int]string)
	for _, item := range strings.Split(args, ",") {
		intKey, err := strconv.Atoi(item)
		if err == nil {
			data[intKey] = item
		}

	}
	return
}
