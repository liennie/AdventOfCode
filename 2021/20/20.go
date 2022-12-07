package main

import (
	"fmt"
	"math"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type image struct {
	min, max space.Point
	data     map[space.Point]bool
	inf      bool
}

func parse(filename string) (image, []bool) {
	ch := load.File(filename)

	line := <-ch
	iea := make([]bool, len(line))
	for i := 0; i < len(line); i++ {
		if line[i] == '#' {
			iea[i] = true
		}
	}
	<-ch // empty line

	min := space.Point{X: math.MaxInt, Y: math.MaxInt}
	max := space.Point{}
	img := map[space.Point]bool{}

	i := 0
	for line := range ch {
		for j := 0; j < len(line); j++ {
			if line[j] == '#' {
				img[space.Point{X: j, Y: i}] = true

				min.X = ints.Min(min.X, j)
				min.Y = ints.Min(min.Y, i)
				max.X = ints.Max(max.X, j)
				max.Y = ints.Max(max.Y, i)
			}
		}
		i++
	}

	return image{
		min:  min,
		max:  max,
		data: img,
	}, iea
}

func step(img image, iea []bool) image {
	res := image{
		min:  img.min,
		max:  img.max,
		data: map[space.Point]bool{},
		inf:  img.inf != iea[0],
	}

	for y := img.min.Y - 1; y <= img.max.Y+1; y++ {
		for x := img.min.X - 1; x <= img.max.X+1; x++ {
			i := 0
			for yd := -1; yd <= 1; yd++ {
				for xd := -1; xd <= 1; xd++ {
					i <<= 1

					p := space.Point{
						X: x + xd,
						Y: y + yd,
					}
					if img.data[p] || (img.inf && (p.X < img.min.X || p.Y < img.min.Y || p.X > img.max.X || p.Y > img.max.Y)) {
						i |= 1
					}
				}
			}

			if iea[i] {
				res.data[space.Point{X: x, Y: y}] = true
				res.min.X = ints.Min(res.min.X, x)
				res.min.Y = ints.Min(res.min.Y, y)
				res.max.X = ints.Max(res.max.X, x)
				res.max.Y = ints.Max(res.max.Y, y)
			}
		}
	}

	return res
}

func (img image) print() {
	for y := img.min.Y; y <= img.max.Y; y++ {
		for x := img.min.X; x <= img.max.X; x++ {
			if img.data[space.Point{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	img, iea := parse(filename)

	// Part 1
	for i := 0; i < 2; i++ {
		img = step(img, iea)
	}
	log.Part1(len(img.data))

	// Part 2
	for i := 2; i < 50; i++ {
		img = step(img, iea)
	}
	log.Part2(len(img.data))
}
