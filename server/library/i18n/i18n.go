package i18n

import "fmt"

func T(key string, args ...interface{}) string {
	s := key
	if 0 < len(args) {
		return fmt.Sprintf(s, args...)
	}
	return s
}
