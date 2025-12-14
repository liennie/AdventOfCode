package main

import (
	"fmt"
	_ "net/http/pprof"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/seq"
)

type Present struct {
	shapes [8][][]bool
	area   int
}

type Region struct {
	w, h     int
	presents []int
}

func flipShape(shape [][]bool) [][]bool {
	shape = cpState(shape)
	slices.Reverse(shape)
	return shape
}

func rotateShape(shape [][]bool, c int) [][]bool {
	shape = cpState(shape)
	for range c {
		for y := 0; y < len(shape)/2; y++ {
			for x := 0; x < (len(shape[y])+1)/2; x++ {
				tmp := shape[y][x]
				shape[y][x] = shape[x][len(shape[y])-y-1]
				shape[x][len(shape[y])-y-1] = shape[len(shape)-y-1][len(shape[y])-x-1]
				shape[len(shape)-y-1][len(shape[y])-x-1] = shape[len(shape)-x-1][y]
				shape[len(shape)-x-1][y] = tmp
			}
		}
	}
	return shape
}

func parse(filename string) ([]Present, []Region) {
	ch := load.Blocks(filename)
	defer channel.Drain(ch)

	presents := []Present{}
	for range 6 {
		block := <-ch
		<-block

		p := Present{}
		shape := [][]bool{}
		for line := range block {
			shape = append(shape, slices.Collect(seq.Map(slices.Values(strings.Split(line, "")), func(s string) bool { return s == "#" })))
		}

		flipped := flipShape(shape)

		p.shapes[0] = shape
		p.shapes[1] = rotateShape(shape, 1)
		p.shapes[2] = rotateShape(shape, 2)
		p.shapes[3] = rotateShape(shape, 3)
		p.shapes[4] = flipped
		p.shapes[5] = rotateShape(flipped, 1)
		p.shapes[6] = rotateShape(flipped, 2)
		p.shapes[7] = rotateShape(flipped, 3)

		for _, row := range shape {
			for _, s := range row {
				if s {
					p.area++
				}
			}
		}

		presents = append(presents, p)
	}

	regions := []Region{}
	re := regexp.MustCompile(`^(\d+)x(\d+): ([\d ]+)$`)
	for line := range <-ch {
		match := re.FindStringSubmatch(line)
		evil.Assert(match != nil, "line not matched %q", line)

		regions = append(regions, Region{
			w:        evil.Atoi(match[1]),
			h:        evil.Atoi(match[2]),
			presents: evil.Fields(match[3]),
		})
	}

	return presents, regions
}

func mkState(w, h int) [][]bool {
	mem := make([]bool, w*h)
	state := [][]bool{}
	for i := range h {
		state = append(state, mem[i*w:(i+1)*w])
	}
	return state
}

func cpState(state [][]bool) [][]bool {
	newState := mkState(len(state[0]), len(state))
	for y := range state {
		copy(newState[y], state[y])
	}
	return newState
}

func printState(state [][]bool, maxx []int) {
	for y, row := range state {
		for _, s := range row {
			if s {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println(" ", maxx[y])
	}
}

func canPlace(state [][]bool, shape [][]bool, x, y int) bool {
	for yy, row := range shape {
		for xx, s := range row {
			if s && state[y+yy][x+xx] {
				return false
			}
		}
	}
	return true
}

func place(state [][]bool, maxx []int, shape [][]bool, x, y int) ([][]bool, []int) {
	newState := cpState(state)
	for yy, row := range shape {
		for xx, s := range row {
			if s {
				newState[y+yy][x+xx] = true
			}
		}
	}
	newMaxx := make([]int, len(maxx))
	copy(newMaxx, maxx)
	for yy := 0; yy < len(shape); yy++ {
		for xx := len(shape[yy]) - 1; xx >= 0; xx-- {
			if shape[yy][xx] {
				newMaxx[y+yy] = ints.Max(newMaxx[y+yy], x+xx)
				break
			}
		}
	}
	return newState, newMaxx
}

func score(state [][]bool, maxx []int, presents []Present, cnts []int) int {
	score := 0
	for p, present := range presents {
		for _, shape := range present.shapes {
			for y := 0; y <= len(state)-len(shape); y++ {
				topx := len(state[y]) - len(shape[0])
				thresh := ints.Max(maxx[y : y+len(shape)]...)
				for x := ints.Min(maxx[y : y+len(shape)]...); x <= topx; x++ {
					if x > thresh {
						score += cnts[p] * (topx - x + 1)
						break
					}

					if canPlace(state, shape, x, y) {
						score += cnts[p]
					}
				}
			}
		}
	}
	return score
}

func canFit(state [][]bool, maxx []int, presents []Present, cnts []int, debug string) bool {
	if ints.Sum(cnts...) == 0 {
		return true
	}

	type top struct {
		present int
		state   [][]bool
		maxx    []int
	}
	max := -1
	var maxInfo top

	for p, present := range presents {
		if cnts[p] == 0 {
			continue
		}
		cnts[p]--

		for _, shape := range present.shapes {
			for y := 0; y <= len(state)-len(shape); y++ {
				// for x := 0; x <= len(state[y])-len(shape[0]); x++ {
				for x := ints.Min(maxx[y : y+len(shape)]...); x <= len(state[y])-len(shape[0]); x++ {
					if !canPlace(state, shape, x, y) {
						continue
					}

					pstate, pmaxx := place(state, maxx, shape, x, y)
					sc := score(pstate, pmaxx, presents, cnts)
					if sc > max {
						max = sc
						maxInfo = top{
							present: p,
							state:   pstate,
							maxx:    pmaxx,
						}
					}
					break
				}
			}
		}

		cnts[p]++
	}

	if max > -1 {
		cnts[maxInfo.present]--
		fmt.Print("\033[H")
		log.Print(debug)
		fmt.Print("\033[2K")
		log.Print(max, cnts)
		printState(maxInfo.state, maxx)
		return canFit(maxInfo.state, maxInfo.maxx, presents, cnts, debug)
	}
	return false
}

func (r Region) canFit(presents []Present, debug string) bool {
	totalArea := 0
	for p, present := range presents {
		totalArea += present.area * r.presents[p]
	}
	if totalArea > r.w*r.h {
		return false
	}

	// don't worry about it
	return true

	fmt.Print("\033[2J")
	return canFit(mkState(r.w, r.h), make([]int, r.h), presents, r.presents, debug)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	presents, regions := parse(filename)

	// Part 1
	st := time.Now()
	cnt := 0
	for r, region := range regions {
		evil.Assert(len(presents) == len(region.presents), "lens don't match %d != %d", len(presents), len(region.presents))
		debug := fmt.Sprintf("region %d / %d   (T: %d; F: %d)   st: %v   run: %v", r+1, len(regions), cnt, r-cnt, st.Format("15:04:05.000000"), time.Since(st))
		cf := region.canFit(presents, debug)
		log.Printf("region %d can fit %v", r+1, cf)
		if cf {
			cnt++
		}
	}
	log.Part1(cnt)
}
