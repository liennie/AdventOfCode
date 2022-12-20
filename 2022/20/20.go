package main

import (
	"container/ring"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func mix(numbers []int, key int, count int) int {
	numberRing := ring.New(len(numbers))
	order := make([]*ring.Ring, 0, len(numbers))
	var zero *ring.Ring
	for _, n := range numbers {
		numberRing.Value = n * key
		order = append(order, numberRing)
		if n == 0 {
			zero = numberRing
		}
		numberRing = numberRing.Next()
	}

	for c := 0; c < count; c++ {
		for _, r := range order {
			val := r.Value.(int) % (len(numbers) - 1)
			if val == 0 {
				continue
			}

			nr := r
			if val < 0 {
				nr = nr.Prev()
				for i := 0; i > val; i-- {
					nr = nr.Prev()
				}
			} else {
				for i := 0; i < val; i++ {
					nr = nr.Next()
				}
			}

			nr.Link(r.Prev().Unlink(1))
		}
	}

	sum := 0
	for r, i := zero, 0; i <= 3000; r, i = r.Next(), i+1 {
		if i == 1000 || i == 2000 || i == 3000 {
			sum += r.Value.(int)
		}
	}

	return sum
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	numbers := evil.SliceAtoi(load.Slice(filename))

	// Part 1
	log.Part1(mix(numbers, 1, 1))

	// Part 2
	log.Part2(mix(numbers, 811589153, 10))
}
