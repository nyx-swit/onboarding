package env

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"time"
)

var _config *viper.Viper

func init() {
	_config = viper.New()
	_config.AutomaticEnv()
	_config.SetDefault("SWIT_RUN_MODE", "debug")
	if _config.GetString("SWIT_RUN_MODE") == "debug" {
		res, err := gotenv.StrictParse(bytes.NewBuffer(ENV))
		if err != nil {
			panic(err)
		}
		for k, v := range res {
			_config.SetDefault(k, v)
		}
	}
}

//file path env.go
func MergeConfig(path string) error {
	_config.WatchConfig()
	_config.SetConfigFile(path)
	return _config.ReadInConfig()
}

func GetString(key string) string {
	return _config.GetString(key)
}

func Get(key string) interface{} {
	return _config.Get(key)
}

func AllKeys() []string {
	return _config.AllKeys()
}

func IsSet(key string) bool {
	return _config.IsSet(key)
}

func Set(key string, value interface{}) {
	_config.Set(key, value)
}

func GetStringSlice(key string) []string {
	return _config.GetStringSlice(key)
}

func GetUint(key string) uint {
	return _config.GetUint(key)
}

func GetUint32(key string) uint32 {
	return _config.GetUint32(key)
}

func GetUint64(key string) uint64 {
	return _config.GetUint64(key)
}

func GetTime(key string) time.Time {
	return _config.GetTime(key)
}

func GetStringMapStringSlice(key string) map[string][]string {
	return _config.GetStringMapStringSlice(key)
}

func GetStringMapString(key string) map[string]string {
	return _config.GetStringMapString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return _config.GetStringMap(key)
}

func GetSizeInBytes(key string) uint {
	return _config.GetSizeInBytes(key)
}

func GetIntSlice(key string) []int {
	return _config.GetIntSlice(key)
}

func GetInt(key string) int {
	return _config.GetInt(key)
}

func GetDuration(key string) time.Duration {
	return _config.GetDuration(key)
}

func GetBool(key string) bool {
	return _config.GetBool(key)
}

func GetFloat64(key string) float64 {
	return _config.GetFloat64(key)
}

func GetInt64(key string) int64 {
	return _config.GetInt64(key)
}

func GetInt32(key string) int32 {
	return _config.GetInt32(key)
}
