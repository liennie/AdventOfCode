package load

import "os"

func Filename() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "input.txt"
}
