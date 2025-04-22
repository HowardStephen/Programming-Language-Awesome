package split

import "strings"

func Split(s, sep string) (result []string) {
	index := strings.Index(s, sep)

	if index == 0 {
		s = s[len(sep):]
		index = strings.Index(s, sep)
	}

	for index > -1 {
		result = append(result, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}

	if s != "" {
		result = append(result, s)
	}

	return
}
