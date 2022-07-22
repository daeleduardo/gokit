package lib

import(
	"fmt"
	"crypto/md5"
)

func GetMd5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}