package auth

import "time"

// keyExpired
func keyExpired(expires int64) bool {
	if expires >= 1 {
		return time.Now().After(time.Unix(expires, 0))
	}
	return false
}
