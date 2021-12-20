package main

import (
	"fmt"
	"math"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type image struct {
	min, max util.Point
	data     map[util.Point]bool
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

	min := util.Point{X: math.MaxInt, Y: math.MaxInt}
	max := util.Point{}
	img := map[util.Point]bool{}

	i := 0
	for line := range ch {
		for j := 0; j < len(line); j++ {
			if line[j] == '#' {
				img[util.Point{X: j, Y: i}] = true

				min.X = util.Min(min.X, j)
				min.Y = util.Min(min.Y, i)
				max.X = util.Max(max.X, j)
				max.Y = util.Max(max.Y, i)
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
		data: map[util.Point]bool{},
		inf:  img.inf != iea[0],
	}

	for y := img.min.Y - 1; y <= img.max.Y+1; y++ {
		for x := img.min.X - 1; x <= img.max.X+1; x++ {
			i := 0
			for yd := -1; yd <= 1; yd++ {
				for xd := -1; xd <= 1; xd++ {
					i <<= 1

					p := util.Point{
						X: x + xd,
						Y: y + yd,
					}
					if img.data[p] || (img.inf && (p.X < img.min.X || p.Y < img.min.Y || p.X > img.max.X || p.Y > img.max.Y)) {
						i |= 1
					}
				}
			}

			if iea[i] {
				res.data[util.Point{X: x, Y: y}] = true
				res.min.X = util.Min(res.min.X, x)
				res.min.Y = util.Min(res.min.Y, y)
				res.max.X = util.Max(res.max.X, x)
				res.max.Y = util.Max(res.max.Y, y)
			}
		}
	}

	return res
}

func (img image) print() {
	for y := img.min.Y; y <= img.max.Y; y++ {
		for x := img.min.X; x <= img.max.X; x++ {
			if img.data[util.Point{X: x, Y: y}] {
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
	defer util.Recover(log.Err)

	const filename = "input.txt"

	img, iea := parse(filename)

	// img.print()

	// Part 1
	for i := 0; i < 2; i++ {
		img = step(img, iea)
		// img.print()
	}
	log.Part1(len(img.data))
}
