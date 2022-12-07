package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type node interface {
	isDir() bool
	size() int
	content() map[string]node
}

type dir struct {
	c map[string]node
}

func (d *dir) isDir() bool {
	return true
}

func (d *dir) size() int {
	sum := 0
	for _, n := range d.c {
		sum += n.size()
	}
	return sum
}

func (d *dir) content() map[string]node {
	return d.c
}

type file struct {
	s int
}

func (f *file) isDir() bool {
	return false
}

func (f *file) size() int {
	return f.s
}

func (f *file) content() map[string]node {
	return nil
}

func foreach(n node, f func(node)) {
	f(n)
	for _, ch := range n.content() {
		foreach(ch, f)
	}
}

func parse(filename string) *dir {
	root := &dir{c: map[string]node{}}
	pwd := []*dir{root}
	lastCmd := ""

	for line := range load.File(filename) {
		if strings.HasPrefix(line, "$") {
			cmd, args, _ := strings.Cut(strings.TrimSpace(strings.TrimPrefix(line, "$")), " ")
			switch cmd {
			case "cd":
				switch args {
				case "..":
					if len(pwd) > 1 {
						pwd = pwd[:len(pwd)-1]
					}
				case "/":
					pwd = pwd[:1]
				default:
					if n, ok := pwd[len(pwd)-1].c[args]; ok {
						if d, ok := n.(*dir); ok {
							pwd = append(pwd, d)
						} else {
							evil.Panic("%s: not a dir", line)
						}
					} else {
						evil.Panic("%s: not found", line)
					}
				}

			case "ls":
				// nothin

			default:
				evil.Panic("unknown cmd %s", cmd)
			}
			lastCmd = cmd
		} else {
			switch lastCmd {
			case "ls":
				size, name, ok := strings.Cut(line, " ")
				if !ok {
					break
				}

				if size == "dir" {
					pwd[len(pwd)-1].c[name] = &dir{c: map[string]node{}}
				} else {
					pwd[len(pwd)-1].c[name] = &file{s: evil.Atoi(size)}
				}
			}
		}
	}

	return root
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	root := parse(filename)

	// Part 1
	sum := 0
	foreach(root, func(n node) {
		if n.isDir() {
			if size := n.size(); size <= 100000 {
				sum += size
			}
		}
	})
	log.Part1(sum)

	// Part 2
	const total = 70000000
	const needed = 30000000

	threshold := needed - (total - root.size())
	if threshold <= 0 {
		evil.Panic("no need to free up space")
	}

	min := total
	foreach(root, func(n node) {
		if n.isDir() {
			if size := n.size(); size >= threshold && size < min {
				min = size
			}
		}
	})
	log.Part2(min)
}
