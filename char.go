package persian

import "strings"

var (
	f2eData = map[string][]string{
		"تَ": {"a", "a"},
		"تِ": {"e", "e"},
		"تُ": {"o", "o"},
		"ﻁْ": {"."},
		"ا":  {"a", "a"},
		"آ":  {"a", "a"},
		"ئ":  {"a"},
		"ء":  {"a"},
		"ب":  {"b"},
		"پ":  {"p"},
		"ت":  {"t"},
		"ث":  {"s"},
		"ج":  {"j"},
		"چ":  {"ch"},
		"ح":  {"h"},
		"خ":  {"kh"},
		"د":  {"d"},
		"ذ":  {"z"},
		"ر":  {"r"},
		"ز":  {"z"},
		"ژ":  {"zh"},
		"س":  {"s"},
		"ش":  {"sh"},
		"ص":  {"s"},
		"ض":  {"z"},
		"ط":  {"t"},
		"ظ":  {"z"},
		"ع":  {"", "a"},
		"غ":  {"gh"},
		"ف":  {"f"},
		"ق":  {"gh"},
		"ک":  {"k"},
		"گ":  {"g"},
		"ل":  {"l"},
		"م":  {"m"},
		"ن":  {"n"},
		"و":  {"v", "o"},
		"ه":  {"h"},
		"ی":  {"y", "i"},
		"ي":  {"y", "i"},
	}
)

func F2E(s string) string {
	var (
		result = ""
		state  = 0
	)
	splited := strings.Split(s, "")
	for i := 0; i < len(splited); i++ {
		ch := splited[i]
		if ch == " " || ch == "." {
			state = 0
			result += " "
			continue
		}
		part, ok := f2eData[ch]
		if !ok {
			result += F2ENum(ch)
			continue
		}
		switch state {
		case 0:
			state = 1
			result += part[0]
		case 1:
			if len(part) > 1 {
				state = 2
				result += part[1]
				continue
			}
			state = 2
			result += "a"
			i--
		case 2:
			state = 3
			result += part[0]
		case 3:
			if len(part) > 1 {
				state = 2
				result += part[1]
				continue
			}
			state = 4
			result += part[0]
		case 5:

		default:
			result += part[0]
			if len(part) > 1 {
				state = 2
				result += part[1]
				continue
			}
			state = 1
			result += part[0]
		}
	}

	result = strings.Replace(result, "aa", "a", -1)
	result = strings.Replace(result, "ao", "o", -1)
	result = strings.Replace(result, "ae", "e", -1)
	result = F2ENum(result)
	return result
}
