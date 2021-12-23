package main

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type burrow struct {
	hallway  [7]byte
	rooms    [4][4]byte
	maxDepth int
}

func (b burrow) String() string {
	c := func(b byte) byte {
		if b == 0 {
			return '.'
		}
		return b
	}

	sb := &strings.Builder{}
	fmt.Fprintln(sb, "#############")
	fmt.Fprintf(sb, "#%c%c.%c.%c.%c.%c%c#\n", c(b.hallway[0]), c(b.hallway[1]), c(b.hallway[2]), c(b.hallway[3]), c(b.hallway[4]), c(b.hallway[5]), c(b.hallway[6]))
	for depth := 0; depth <= b.maxDepth; depth++ {
		fmt.Fprintf(sb, "###%c#%c#%c#%c###\n", c(b.rooms[0][depth]), c(b.rooms[1][depth]), c(b.rooms[2][depth]), c(b.rooms[3][depth]))
	}
	fmt.Fprintln(sb, "  #########")
	return sb.String()
}

func roomIndex(r, d int) int {
	return 7 + r + d*4
}

func room(i int) (int, int) {
	return (i - 7) % 4, (i - 7) / 4
}

func (b *burrow) get(i int) *byte {
	if i < 7 {
		return &b.hallway[i]
	}
	r, d := room(i)
	return &b.rooms[r][d]
}

func (b burrow) dist(hall, roomI int) int {
	room, depth := room(roomI)
	roomHall := room + 1
	if hall > roomHall {
		roomHall++
	}

	min, max := util.Min(roomHall, hall), util.Max(roomHall, hall)

	return (util.Clamp(max, 0, 1) - util.Clamp(min, 0, 1)) +
		((util.Clamp(max, 1, 5) - util.Clamp(min, 1, 5)) * 2) +
		(util.Clamp(max, 5, 6) - util.Clamp(min, 5, 6)) +
		depth + 2
}

func (b burrow) obstacles(hall, roomI int) []int {
	res := []int{}

	room, depth := room(roomI)
	roomHall := room + 1
	if hall <= roomHall {
		roomHall++
	}

	for d := 0; d < depth; d++ {
		res = append(res, roomIndex(room, d))
	}
	min, max := util.Min(roomHall, hall), util.Max(roomHall, hall)
	for h := min + 1; h <= max-1; h++ {
		res = append(res, h)
	}

	return res
}

var cost = map[byte]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

func (b burrow) move(from, to int) (burrow, int, bool) {
	hall := util.Min(from, to)
	roomI := util.Max(from, to)

	fromAmph := *b.get(from)
	toAmph := *b.get(to)

	for _, obstacle := range b.obstacles(hall, roomI) {
		if *b.get(obstacle) != 0 {
			return b, 0, false
		}
	}

	if to >= 7 {
		toRoom, toDepth := room(to)
		amph := fromAmph
		if toRoom != int(amph-'A') {
			return b, 0, false
		}
		for depth := b.maxDepth; depth > toDepth; depth-- {
			if b.rooms[toRoom][depth] != amph {
				return b, 0, false
			}
		}
	}

	cost := cost[fromAmph] * b.dist(hall, roomI)
	*b.get(from), *b.get(to) = toAmph, fromAmph
	return b, cost, true
}

var organized2 = burrow{
	rooms: [4][4]byte{
		{'A', 'A'},
		{'B', 'B'},
		{'C', 'C'},
		{'D', 'D'},
	},
	maxDepth: 1,
}

var organized4 = burrow{
	rooms: [4][4]byte{
		{'A', 'A', 'A', 'A'},
		{'B', 'B', 'B', 'B'},
		{'C', 'C', 'C', 'C'},
		{'D', 'D', 'D', 'D'},
	},
	maxDepth: 3,
}

func (b burrow) isOrganized() bool {
	return b == organized2 || b == organized4
}

func parse(filename string) burrow {
	ch := load.File(filename)
	defer util.Drain(ch)

	res := burrow{
		maxDepth: 1,
	}

	<-ch // #############
	<-ch // #...........#

	p := strings.Split(<-ch, "#")
	for i := 3; i <= 6; i++ {
		res.rooms[i-3][0] = p[i][0]
	}
	p = strings.Split(<-ch, "#")
	for i := 1; i <= 4; i++ {
		res.rooms[i-1][1] = p[i][0]
	}

	return res
}

type move struct {
	newBurrow burrow
	cost      int
}

func generateMoves(b burrow) []move {
	moves := []move{}

	for hall := 0; hall < 7; hall++ {
		if b.hallway[hall] != 0 {
			for room := 0; room < 4; room++ {
				for depth := b.maxDepth; depth >= 0; depth-- {
					if b.rooms[room][depth] == 0 {
						if nb, c, ok := b.move(hall, roomIndex(room, depth)); ok {
							moves = append(moves, move{
								newBurrow: nb,
								cost:      c,
							})
						}
						break
					}
				}
			}
		}
	}

	for room := 0; room < 4; room++ {
		for depth := 0; depth <= b.maxDepth; depth++ {
			if b.rooms[room][depth] != 0 {
				for hall := 0; hall < 7; hall++ {
					if b.hallway[hall] == 0 {
						if nb, c, ok := b.move(roomIndex(room, depth), hall); ok {
							moves = append(moves, move{
								newBurrow: nb,
								cost:      c,
							})
						}
					}
				}
				break
			}
		}
	}

	return moves
}

func organize(b burrow) int {
	cost := map[burrow]int{
		b: 0,
	}
	next := list.New()
	next.PushBack(move{
		newBurrow: b,
		cost:      0,
	})

	for next.Len() > 0 {
		minCost := next.Front().Value.(move).cost
		minBurrow := next.Front().Value.(move).newBurrow

		if minBurrow.isOrganized() {
			return minCost
		}

		next.Remove(next.Front())

		for _, m := range generateMoves(minBurrow) {
			oc, ok := cost[m.newBurrow]
			if c := m.cost + minCost; !ok || c < oc {
				cost[m.newBurrow] = c

				nm := move{
					newBurrow: m.newBurrow,
					cost:      c,
				}

				ins := false
				for el := next.Front(); el != nil; el = el.Next() {
					if c < el.Value.(move).cost {
						next.InsertBefore(nm, el)
						ins = true
						break
					}
				}
				if !ins {
					next.PushBack(nm)
				}
			}
		}
	}

	util.Panic("No solution found")
	return 0
}

func unfold(b burrow) burrow {
	return burrow{
		hallway: b.hallway,
		rooms: [4][4]byte{
			{b.rooms[0][0], 'D', 'D', b.rooms[0][1]},
			{b.rooms[1][0], 'C', 'B', b.rooms[1][1]},
			{b.rooms[2][0], 'B', 'A', b.rooms[2][1]},
			{b.rooms[3][0], 'A', 'C', b.rooms[3][1]},
		},
		maxDepth: 3,
	}
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	b := parse(filename)

	log.Println("Start")

	// Part 1
	log.Part1(organize(b))

	// Part 2
	log.Part2(organize(unfold(b)))
}
