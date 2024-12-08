package load

import (
	"bufio"
	"container/list"
	"io"
	"iter"
	"os"
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func lines(filename string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(filename)
		if err != nil {
			evil.Panic("os.Open: %w", err)
		}
		defer file.Close()

		r := bufio.NewReader(file)
		for {
			l, err := r.ReadString('\n')
			if len(l) > 0 {
				if !yield(strings.TrimSuffix(l, "\n")) {
					return
				}
			}
			if err != nil {
				if err != io.EOF {
					evil.Panic("ReadString: %w", err)
				}
				return
			}
		}
	}
}

func File(filename string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		for line := range lines(filename) {
			ch <- line
		}
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

		for line := range lines(filename) {
			if line == "" {
				if bch != nil {
					close(bch)
					bch = nil
				}
				blank = true
				continue
			}

			if blank {
				bch = make(chan string)
				ch <- bch
				blank = false
			}
			bch <- line
		}
	}()

	return ch
}

func Slice(filename string) []string {
	return slices.Collect(lines(filename))
}

func Set(filename string) set.String {
	return set.Collect(lines(filename))
}

func List(filename string) *list.List {
	res := list.New()

	for line := range lines(filename) {
		res.PushBack(line)
	}

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

func Parse[T any](filename string, parse func(line string) T) []T {
	var res []T
	for line := range lines(filename) {
		res = append(res, parse(line))
	}
	return res
}

func Grid(filename string, cb func(p space.Point, r rune)) {
	y := 0
	for line := range lines(filename) {
		for x, r := range line {
			cb(space.Point{X: x, Y: y}, r)
		}
		y++
	}
}
