package nasType

import (
    "fmt"
    "strings"
)

func GetHexString(bytes []uint8, sep string) string {
	if len(bytes) > 0 {

		hexvals := make([]string, 0, len(bytes))
		for idx := range bytes {
			chr := fmt.Sprintf("%.2x", bytes[idx])
			if chr != "" {
				hexvals = append(hexvals, chr)
			}
		}
		return strings.Join(hexvals, sep)
	}

	return ""
}
