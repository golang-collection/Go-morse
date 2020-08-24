package tools

import (
	"errors"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-24 10:22
* @Description:
**/

var morseMap = map[byte]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",
}

//根据字符串生成摩尔斯密码
func GenerateMorse(str string) (string, error) {
	length := len(str)
	if length == 0{
		return "", errors.New("length must > 1")
	}
	values := make([]string, 0)
	for _, v := range str{
		if value, ok := morseMap[byte(v)]; ok{
			values = append(values, value)
		}
	}
	result := strings.Join(values, "")
	return result, nil
}