package env

import (
	"os"
	"strconv"
)

func Get(key string, defaultValue ...string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return func() string {
			if len(defaultValue) != 0 {
				return defaultValue[0]
			}
			return ""
		}()
	}
	return v
}

func GetInt(key string, defaultValue ...int) int {
	v := os.Getenv(key)
	if len(v) == 0 {
		return func() int {
			if len(defaultValue) != 0 {
				return defaultValue[0]
			}
			return 0
		}()
	}
	i, _ := strconv.Atoi(v)
	return i
}
