package crack

import (
	"crypto/sha1"
	"fmt"
)

func Hash(password string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(password)))
}
