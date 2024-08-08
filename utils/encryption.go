package utils

import (
	"encoding/base64"
	"slices"
	"strings"
)

const (
	key = "kjsdbaiufguiassdhu3ihf7843ghowienfu934gbuewhc9293rh9bgu942bgu9"
)

func Encrypt(code string) string {
	res := []byte(code)
	for i := 0; i < len(res); i++ {
		res[i]++
	}
	s1 := string(res)
	slices.Reverse(res)
	s2 := string(res)
	baseS := base64.StdEncoding.EncodeToString([]byte(s2 + key + s1))
	//所有结尾的'='替换为%3D
	return strings.ReplaceAll(baseS, "=", "%3D")
}
func Decrypt(code string) string {
	code = strings.ReplaceAll(code, "%3D", "=")
	strsbytes, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return ""
	}
	strs := strings.Split(string(strsbytes), key)
	if len(strs) != 2 {
		return ""
	}
	res := []byte(strs[1])
	for i := 0; i < len(res); i++ {
		res[i]--
	}
	return string(res)
}
