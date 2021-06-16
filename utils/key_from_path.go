package utils

import (
	"strings"
)

func KeyFromPath(path string, position int) string {
	splitPath := strings.Split(path, "/")
	if len(splitPath) > 2 && splitPath[2] != "" {
		return strings.Split(splitPath[2], "?")[0]
	}
	return ""
}
