package main

import (
	"fmt"
	"strings"
)

func sanitizeIdentifierComponent(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	b.Grow(len(s))
	prevUnderscore := false
	for _, r := range s {
		isAllowed := (r >= 'a' && r <= 'z') ||
			(r >= '0' && r <= '9') ||
			r == '_' || r == '-'
		if !isAllowed {
			if !prevUnderscore {
				b.WriteRune('_')
				prevUnderscore = true
			}
			continue
		}
		if r == '_' {
			if prevUnderscore {
				continue
			}
			prevUnderscore = true
		} else {
			prevUnderscore = false
		}
		b.WriteRune(r)
	}
	result := strings.Trim(b.String(), "_")
	if result == "" {
		result = "unnamed"
	}
	return result
}
func main() {
	fmt.Println(sanitizeIdentifierComponent("hdn-server"))
}
