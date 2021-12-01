package load

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func File(filename string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		file, err := os.Open(filename)
		if err != nil {
			panic(fmt.Errorf("os.Open: %w", err))
		}
		defer file.Close()

		r := bufio.NewReader(file)
		for {
			l, err := r.ReadString('\n')
			if len(l) > 0 {
				ch <- strings.TrimSuffix(l, "\n")
			}
			if err != nil {
				if err != io.EOF {
					panic(fmt.Errorf("ReadString: %w", err))
				}
				return
			}
		}
	}()

	return ch
}
