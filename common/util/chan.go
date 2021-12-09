package util

func Drain(ch <-chan string) {
	for range <-ch {
	}
}
