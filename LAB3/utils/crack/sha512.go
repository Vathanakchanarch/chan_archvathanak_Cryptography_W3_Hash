package crack

import (
	"crypto/sha512"
	"fmt"
)

func Hash(password string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(password)))
}
