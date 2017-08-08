package env

import "os"

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
