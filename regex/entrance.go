package regex

import "regexp"

type Entity struct {
	regex map[string]string
}

func New() *Entity {
	result := &Entity{}
	result.regex = make(map[string]string)
	return result
}

func (s *Entity) Init() {
	s.regex["phone"] = `^\+\d{1,3}$` // 校验是否是手机号

	s.regex["+86"] = `^1[3-9]\d{9}$`                 // 中国大陆，手机号码通常是11位，以 13 - 19 开头
	s.regex["+852"] = `^\d{8}$`                      // 中国香港，通常8位数字，但注意实际格式可能因运营商而异
	s.regex["+853"] = `^\d{8}$`                      // 中国澳门，通常8位数字，但注意实际格式可能因运营商而异
	s.regex["+886"] = `^(09[0-9]{8}|[289][0-9]{9})$` // 中国台湾，这里提供一个可能的格式，但可能不完整

	s.regex["email"] = `^[^\s@]+@[^\s@]+\.[^\s@]+$`             // 邮箱
	s.regex["ipv4"] = `(?m)^(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})$` // ipv4
}

func (s *Entity) Append(key string, reg string) {
	s.regex[key] = reg
}

func (s *Entity) Validate(key string, val string) bool {
	if str, ok := s.regex[key]; ok {
		compile, err := regexp.Compile(str)
		if err != nil {
			return false
		}
		return compile.MatchString(val)
	}
	return false
}

func (s *Entity) Get(key string, val string) string {
	if str, ok := s.regex[key]; ok {
		compile := regexp.MustCompile(str)
		return compile.FindString(val)
	}
	return ""
}

func (s *Entity) GetAll(key string, val string) []string {
	if str, ok := s.regex[key]; ok {
		compile := regexp.MustCompile(str)
		return compile.FindAllString(val, -1)
	}
	return nil
}
