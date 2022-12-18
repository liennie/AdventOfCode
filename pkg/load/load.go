package load

import (
	"bufio"
	"container/list"
	"io"
	"os"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/set"
)

func callback(filename string, f func(line string)) {
	file, err := os.Open(filename)
	if err != nil {
		evil.Panic("os.Open: %w", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	for {
		l, err := r.ReadString('\n')
		if len(l) > 0 {
			f(strings.TrimSuffix(l, "\n"))
		}
		if err != nil {
			if err != io.EOF {
				evil.Panic("ReadString: %w", err)
			}
			return
		}
	}
}

func File(filename string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		callback(filename, func(line string) {
			ch <- line
		})
	}()

	return ch
}

func Blocks(filename string) <-chan (<-chan string) {
	ch := make(chan (<-chan string))

	go func() {
		defer close(ch)

		blank := true
		var bch chan string

		defer func() {
			if bch != nil {
				close(bch)
			}
		}()

		callback(filename, func(line string) {
			if line == "" {
				if bch != nil {
					close(bch)
					bch = nil
				}
				blank = true
				return
			}

			if blank {
				bch = make(chan string)
				ch <- bch
				blank = false
			}
			bch <- line
		})
	}()

	return ch
}

func Slice(filename string) []string {
	res := []string{}

	callback(filename, func(line string) {
		res = append(res, line)
	})

	return res
}

func Set(filename string) set.String {
	res := set.String{}

	callback(filename, func(line string) {
		res.Add(line)
	})

	return res
}

func List(filename string) *list.List {
	res := list.New()

	callback(filename, func(line string) {
		res.PushBack(line)
	})

	return res
}

func Line(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		evil.Panic("os.Open: %w", err)
	}
	defer file.Close()

	l, err := bufio.NewReader(file).ReadString('\n')
	if err != nil && err != io.EOF {
		evil.Panic("ReadString: %w", err)
	}

	return strings.TrimSuffix(l, "\n")
}
