package persian

import (
	"strconv"
	"strings"
)

var (
	f2eNumData = map[string]string{
		"۱": "1",
		"۲": "2",
		"۳": "3",
		"۴": "4",
		"۵": "5",
		"۶": "6",
		"۷": "7",
		"۸": "8",
		"۹": "9",
		"۰": "0",
	}
)

func F2Int(s string) int {
	s = strings.Replace(s, " ", "", -1)
	s = F2ENum(s)
	num, _ := strconv.Atoi(s)
	return num
}

func F2Float(s string) float64 {
	s = strings.Replace(s, " ", "", -1)
	s = F2ENum(s)
	num, _ := strconv.ParseFloat(s, 64)
	return num
}

func F2ENum(s string) string {
	result := ""
	for _, rune := range s {
		char := string(rune)
		num, ok := f2eNumData[char]
		if !ok {
			result += char
			continue
		}
		result += num
	}
	return result
}
