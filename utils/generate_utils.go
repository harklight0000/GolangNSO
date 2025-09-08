package utils

import "strings"

func ParseFlags(args []string) map[string]string {
	flags := make(map[string]string)
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if strings.Contains(arg, "=") {
				parts := strings.Split(arg, "=")
				flags[parts[0][1:]] = parts[1]
			} else {
				flags[arg] = ""
			}
		}
	}
	return flags
}
