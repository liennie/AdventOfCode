package main

import (
	"slices"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []int {
	return evil.Split(load.Line(filename), "")
}

func createMemory(blocks []int) []int {
	memory := make([]int, 0, ints.Sum(blocks...))
	for i, n := range blocks {
		isFile := i%2 == 0
		id := -1
		if isFile {
			id = i / 2
		}

		for range n {
			memory = append(memory, id)
		}
	}
	return memory
}

func memoryChecksum(memory []int) int {
	total := 0
	for i, n := range memory {
		if n == -1 {
			continue
		}

		total += i * n
	}
	return total
}

type File struct {
	id  int
	len int
}

func createFiles(blocks []int) []File {
	files := make([]File, len(blocks))
	for i, n := range blocks {
		isFile := i%2 == 0
		id := -1
		if isFile {
			id = i / 2
		}

		files[i] = File{
			id:  id,
			len: n,
		}
	}
	return files
}

func fileChecksum(files []File) int {
	total := 0
	i := 0
	for _, f := range files {
		if f.id == -1 {
			i += f.len
			continue
		}

		for range f.len {
			total += i * f.id
			i++
		}
	}
	return total
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	blocks := parse(filename)

	// Part 1
	memory := createMemory(blocks)
	for i, j := 0, len(memory)-1; i < j; {
		switch {
		case memory[i] != -1:
			i++

		case memory[j] == -1:
			j--

		default:
			// memory[i] == -1
			// memory[j] != -1

			memory[i], memory[j] = memory[j], memory[i]
		}
	}
	log.Part1(memoryChecksum(memory))

	// Part 2
	files := createFiles(blocks)
	for i, j := 0, len(files)-1; i < j; {
		switch {
		case files[i].id != -1:
			i++

		case files[j].id == -1:
			j--

		default:
			// files[i].id == -1
			// files[j].id != -1
			file := &files[j]
		free:
			for k := i; k < j; k++ {
				free := &files[k]
				if free.id != -1 {
					continue
				}

				switch {
				case free.len < file.len:
					continue

				case free.len == file.len:
					free.id = file.id
					file.id = -1

					break free

				case free.len > file.len:
					fileCpy := *file

					free.len -= file.len
					file.id = -1

					files = slices.Insert(files, k, fileCpy)
					j++

					break free
				}
			}
			j--
		}
	}
	log.Part2(fileChecksum(files))
}
