package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// changeEventHandle 配置变更处理器
var changeEventHandle []func(e fsnotify.Event)
var eventLock sync.Mutex
var once sync.Once

// CfgEnv 配置文件路径,允许在初始化前,由外部包赋值  debug / test / release by gin
var CfgEnv string

var (
	//go:embed debug_config.yaml
	debugConfigs []byte

	//go:embed test_config.yaml
	testConfigs []byte

	//go:embed release_config.yaml
	releaseConfigs []byte
)

func Init() {
	once.Do(func() {
		var r io.Reader
		switch CfgEnv {
		case gin.DebugMode:
			r = bytes.NewReader(debugConfigs)
		case gin.TestMode:
			r = bytes.NewReader(testConfigs)
		case gin.ReleaseMode:
			r = bytes.NewReader(releaseConfigs)
		default:
			fmt.Println("Env 参数不合法")
			os.Exit(1)
		}

		// read in environment variables that match
		viper.SetConfigType("yaml")
		viper.SetConfigName(CfgEnv + "_config")
		viper.AddConfigPath("./config/")
		fmt.Println("Env is :", CfgEnv)

		//viper.SetEnvPrefix("FT")
		viper.AutomaticEnv()
		if err := viper.ReadConfig(r); err == nil {
			fmt.Printf("\033[1;30;42m[info]\033[0m using config file %s\n", viper.ConfigFileUsed())
		} else {
			fmt.Printf("\033[1;30;41m[error]\033[0m using config file error %s\n", err.Error())
			os.Exit(1)
		}
	})
}

// IsDevelopment 是否是测试环境
func IsDevelopment() bool {
	if CfgEnv == gin.ReleaseMode {
		return false
	}
	return true
}

// onConfigChange 循环执行事件调用
func onConfigChange(e fsnotify.Event) {
	for _, f := range changeEventHandle {
		f(e)
	}
}

// RegisterChangeEvent 注册配置变更事件
func RegisterChangeEvent(f func(e fsnotify.Event)) {
	eventLock.Lock()
	defer eventLock.Unlock()

	changeEventHandle = append(changeEventHandle, f)
}

// onConfigChange 循环执行事件调用
//func onConfigChange(e fsnotify.Event) {
//	for _, f := range changeEventHandle {
//		f(e)
//	}
//}

// GetBool returns the value associated with the key as a boolean.
func GetBool(k string) bool { return viper.GetBool(k) }

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	val := viper.GetString(key)
	if val == "" {
		fmt.Printf("[ERROR] 请注意：Get config key:%s, val:%s \r\n", key, val)
	}
	return val
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int { return viper.GetInt(key) }

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32 { return viper.GetInt32(key) }

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 { return viper.GetInt64(key) }

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint { return viper.GetUint(key) }

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32 { return viper.GetUint32(key) }

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64 { return viper.GetUint64(key) }

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 { return viper.GetFloat64(key) }

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time { return viper.GetTime(key) }

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration { return viper.GetDuration(key) }

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) []string { return viper.GetStringSlice(key) }

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) map[string]interface{} { return viper.GetStringMap(key) }

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string { return viper.GetStringMapString(key) }

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

// GetSizeInBytes returns the size of the value associated with the given key
// in bytes.
func GetSizeInBytes(key string) uint { return viper.GetSizeInBytes(key) }

func Set(key string, value interface{}) {
	viper.Set(key, value)
}
