package hasher

import (
	"crypto/sha256"
	"fmt"
)

func Hash(input string) string {
	sum := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", sum[:])
}
