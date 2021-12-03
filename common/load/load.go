package load

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/liennie/AdventOfCode/common/set"
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

func Slice(filename string) []string {
	res := []string{}

	for line := range File(filename) {
		res = append(res, line)
	}

	return res
}

func Set(filename string) set.String {
	res := set.String{}
	res.Add(Slice(filename)...)
	return res
}

func List(filename string) *list.List {
	res := list.New()
	for _, item := range Slice(filename) {
		res.PushBack(item)
	}
	return res
}
