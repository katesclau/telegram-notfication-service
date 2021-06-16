package utils

import (
	"fmt"
	"strings"
)

type Stringer interface {
	String() string
}

func BuildString(all ...interface{}) string {
	builder := strings.Builder{}
	for part := range all {
		switch str := all[part].(type) {
		case string:
			builder.WriteString(str)
		case float32:
			builder.WriteString(fmt.Sprintf("%f", str))
		case float64:
			builder.WriteString(fmt.Sprintf("%.2f", str))
		case int:
			builder.WriteString(fmt.Sprintf("%d", str))
		case Stringer:
			builder.WriteString(str.String())
		}

	}
	return builder.String()
}
