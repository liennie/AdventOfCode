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
