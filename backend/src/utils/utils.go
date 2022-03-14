package utils

import "strings"

func ParseValues(values string) (out []string) {
	tokens := strings.Split(values, ",")
	for _, t := range tokens {
		out = append(out, strings.TrimSpace(t))
	}
	return
}
