package tools

import (
	"crypto/md5"
	"encoding/hex"
)

var IP string
var Port string

func Md5Encode(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))

}
func ServerSetting(ipaddress string, p string) {
	IP = ipaddress
	Port = p
}
