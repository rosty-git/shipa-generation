package utils

import "strings"

func ParseValues(values string) (out []string) {
	tokens := strings.Split(values, ",")
	for _, t := range tokens {
		val := strings.TrimSpace(t)
		if val != "" {
			out = append(out, val)
		}
	}
	return
}
