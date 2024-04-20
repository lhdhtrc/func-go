package strings

import "regexp"

func GetIpv4(address string) []string {
	ipv4 := `(?m)^(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})$`
	reg := regexp.MustCompile(ipv4)
	return reg.FindAllString(address, -1)
}
