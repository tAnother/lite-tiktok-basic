package config

var IP string
var Port string

func SetIPAndPort(ipaddress string, p string) {
	IP = ipaddress
	Port = p
}
