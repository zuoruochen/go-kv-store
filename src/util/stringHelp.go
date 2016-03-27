package util

import (
	"strings"
	"errors"
)

func TrimSpace(s string) (ret []string) {
	s = strings.Trim(s, " ")
	i := 0
	j := 0

	for i < len(s) {
		if s[i] == ' ' {
			ret = append(ret, s[j:i])
			i++
			for s[i] == ' ' {
				i++
			}
			j = i
			i++
		} else {
			i++
		}
	}
	ret = append(ret, s[j:i])
	return ret
}

// get a word from a string
// return a word and left string
func GetWord(str string)(string,string,error) {
	s := strings.TrimLeft(str," ")
	if len(s) == 0 {
		return "","",errors.New("blank string")
	}
	sl := strings.Split(s," ")
	word := sl[0]
	var idx = len(word)
	if idx == len(s) {
		return word,"",nil
	}
	return word,s[idx:len(s)],nil
}