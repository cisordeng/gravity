package xenon

import (
	"crypto/md5"
	"fmt"
)

func String2MD5(unencrypted string) (encrypted string) {
	encrypted = fmt.Sprintf("%x", md5.Sum([]byte(unencrypted)))
	return encrypted
}
