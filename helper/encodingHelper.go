package helper

import "github.com/axgle/mahonia"

func EncodingAlUrl(response string) string {
	c := mahonia.NewDecoder("gbk").ConvertString(response)
	return mahonia.NewEncoder("UTF-8").ConvertString(c)
}

func EncodingUTF8(rawString string) string {
	return mahonia.NewDecoder("UTF-8").ConvertString(rawString)
}
