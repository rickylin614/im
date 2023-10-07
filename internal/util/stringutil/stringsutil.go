package stringutil

import "strings"

func Join(parts ...string) string {
	var builder strings.Builder
	for _, part := range parts {
		builder.WriteString(part)
	}
	return builder.String()
}
