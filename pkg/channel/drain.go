package channel

func Drain[T any](ch <-chan T) {
	for range ch {
	}
}