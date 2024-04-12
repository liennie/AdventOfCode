package main

import (
	"cmp"
	"math"
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Hailstone struct {
	pos space.Point3
	vel space.Point3
}

func parsePoint(raw string) space.Point3 {
	n := evil.SplitN(raw, ",", 3)
	evil.Assert(len(n) == 3, "invalid coords")
	return space.Point3{
		X: n[0],
		Y: n[1],
		Z: n[2],
	}
}

func parse(filename string) []Hailstone {
	res := []Hailstone{}
	for line := range load.File(filename) {
		pos, vel, ok := strings.Cut(line, "@")
		evil.Assert(ok, "invalid format")
		res = append(res, Hailstone{
			pos: parsePoint(pos),
			vel: parsePoint(vel),
		})
	}
	return res
}

type Line2 struct {
	k, c float64
}

func newLine2(pos, vel space.Point3) Line2 {
	if vel.X < 0 {
		vel = vel.Flip()
	}

	return Line2{
		k: float64(vel.Y) / float64(vel.X),
		c: float64(pos.Y) - float64(vel.Y)*(float64(pos.X)/float64(vel.X)),
	}
}

func almostEquals(a, b float64) bool {
	return math.Abs(a-b) < math.Pow(10, math.Log10(math.Abs(a))-9)
}

func intersect2(a, b Line2) (float64, float64) {
	if a.k == b.k {
		return math.Inf(1), math.Inf(1)
	}

	x := (b.c - a.c) / (a.k - b.k)
	y := x*a.k + a.c
	y2 := x*b.k + b.c
	evil.Assert(almostEquals(y, y2), a, b, x, y, y2, y-y2)
	return x, y
}

func distanceSquaredAtTime(a, b Hailstone, time float64) float64 {
	ax := float64(a.pos.X) + time*float64(a.vel.X)
	ay := float64(a.pos.Y) + time*float64(a.vel.Y)
	az := float64(a.pos.Z) + time*float64(a.vel.Z)
	bx := float64(b.pos.X) + time*float64(b.vel.X)
	by := float64(b.pos.Y) + time*float64(b.vel.Y)
	bz := float64(b.pos.Z) + time*float64(b.vel.Z)
	dx := ax - bx
	dy := ay - by
	dz := az - bz

	return dx*dx + dy*dy + dz*dz
}

func closestDistanceTime(a, b Hailstone) float64 {
	// p = [x, y, z]
	// v = [dx, dy, dz]

	// distp = (p1 + t * v1) - (p2 + t * v2)
	// distp = p1 - p2 + t * (v1 - v2)
	// distp' = v1 - v2

	// (g(f))' = g'(f) * f'
	// (distp^2)' = 2 * distp * distp'
	// (distp^2)' = 2 * (p1 - p2 + t * (v1 - v2)) * (v1 - v2)
	// (distp^2)' = (2 p1 - 2 p2 + 2 t v1 - 2 t v2) * (v1 - v2)
	// (distp^2)' = 2p1v1 - 2p2v1 + 2tv1^2 - 2tv1v2 - 2p1v2 + 2p2v2 - 2tv1v2 + 2tv2^2
	// (distp^2)' = 2p1v1 - 2p2v1 - 2p1v2 + 2p2v2 + 2tv1^2 - 4tv1v2 + 2tv2^2
	// (distp^2)' = t * (2v1^2 - 4v1v2 + 2v2^2) + 2p1v1 - 2p2v1 - 2p1v2 + 2p2v2

	// dist = distx^2 + disty^2 + distz^2
	// dist' = (distx^2)' + (disty^2)' + (distz^2)'

	// dist'(t) = 0
	//
	// t * (2dx1^2 - 4dx1dx2 + 2dx2^2) + 2x1dx1 - 2x2dx1 - 2x1dx2 + 2x2dx2 +
	// t * (2dy1^2 - 4dy1dy2 + 2dy2^2) + 2y1dy1 - 2y2dy1 - 2y1dy2 + 2y2dy2 +
	// t * (2dz1^2 - 4dz1dz2 + 2dz2^2) + 2z1dz1 - 2z2dz1 - 2z1dz2 + 2z2dz2 = 0
	//
	// t =
	// (-2x1dx1 + 2x2dx1 + 2x1dx2 - 2x2dx2 - 2y1dy1 + 2y2dy1 + 2y1dy2 - 2y2dy2 - 2z1dz1 + 2z2dz1 + 2z1dz2 - 2z2dz2) /
	// (2dx1^2 - 4dx1dx2 + 2dx2^2 + 2dy1^2 - 4dy1dy2 + 2dy2^2 + 2dz1^2 - 4dz1dz2 + 2dz2^2)
	//
	// t =
	// ((x1dx2 + x2dx1 - x1dx1 - x2dx2) + (y1dy2 + y2dy1 - y1dy1 - y2dy2) + (z1dz2 + z2dz1 - z1dz1 - z2dz2)) /
	// ((dx1^2 + dx2^2 - 2dx1dx2) + (dy1^2 + dy2^2 - 2dy1dy2) + (dz1^2 + dz2^2 - 2dz1dz2))

	x1 := float64(a.pos.X)
	y1 := float64(a.pos.Y)
	z1 := float64(a.pos.Z)
	x2 := float64(b.pos.X)
	y2 := float64(b.pos.Y)
	z2 := float64(b.pos.Z)
	dx1 := float64(a.vel.X)
	dy1 := float64(a.vel.Y)
	dz1 := float64(a.vel.Z)
	dx2 := float64(b.vel.X)
	dy2 := float64(b.vel.Y)
	dz2 := float64(b.vel.Z)

	nx := x1*dx2 + x2*dx1 - x1*dx1 - x2*dx2
	ny := y1*dy2 + y2*dy1 - y1*dy1 - y2*dy2
	nz := z1*dz2 + z2*dz1 - z1*dz1 - z2*dz2

	dx := dx1*dx1 + dx2*dx2 - 2*dx1*dx2
	dy := dy1*dy1 + dy2*dy2 - 2*dy1*dy2
	dz := dz1*dz1 + dz2*dz2 - 2*dz1*dz2

	return (nx + ny + nz) / (dx + dy + dz)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	hailstones := parse(filename)

	const from = 200000000000000
	const to = 400000000000000
	// const from = 7
	// const to = 27

	// Part 1
	cnt := 0
	for i := 0; i < len(hailstones)-1; i++ {
		a := hailstones[i]
		for j := i + 1; j < len(hailstones); j++ {
			b := hailstones[j]

			x, y := intersect2(newLine2(a.pos, a.vel), newLine2(b.pos, b.vel))
			if x >= from && x <= to && y >= from && y <= to {
				if ((a.vel.X > 0) == (x > float64(a.pos.X))) && ((b.vel.X > 0) == (x > float64(b.pos.X))) {
					cnt++
				}
			}
		}
	}
	log.Part1(cnt)

	// Part 2
	log.Println("Calculating distances")
	type closestDist struct {
		time float64
		dist float64
		a, b Hailstone
	}
	dists := []closestDist{}
	for i, a := range hailstones {
		for j, b := range hailstones {
			if i == j {
				continue
			}

			t := closestDistanceTime(a, b)
			if t < 0 || math.IsInf(t, 1) || math.IsNaN(t) {
				continue
			}

			dists = append(dists, closestDist{
				time: t,
				dist: distanceSquaredAtTime(a, b, t),
				a:    a,
				b:    b,
			})
		}
	}
	log.Printf("Found %d valid distance pairs", len(dists))
	slices.SortFunc(dists, func(a, b closestDist) int { return cmp.Compare(a.dist, b.dist) })

	totalLoss := func(start Hailstone) float64 {
		total := 0.
		for _, h := range hailstones {
			t := closestDistanceTime(h, start)
			if t < 0 || math.IsInf(t, 1) || math.IsNaN(t) {
				t = 0
			}

			total += distanceSquaredAtTime(h, start, t)
		}
		return total
	}

	minLoss := math.Inf(1)
	min := Hailstone{}
	for _, cd := range dists {
		// log.Printf("Trying pair %d", i+1)

		pairat := int(math.Floor(cd.time))
		pairbt := pairat + 1

		pairMinLoss := math.Inf(1)
		pairMin := Hailstone{}

		for {
			type try struct {
				at, bt int
			}

			tryMinLoss := math.Inf(1)
			tryMin := Hailstone{}
			tryat := pairat
			trybt := pairbt

			for _, tr := range []try{
				{at: pairat, bt: pairbt},

				{at: pairat, bt: pairbt + 1},
				{at: pairat, bt: pairbt + 100},
				{at: pairat, bt: pairbt + 10000},
				{at: pairat, bt: pairbt + 1000000},

				{at: pairat, bt: pairbt - 1},
				{at: pairat, bt: pairbt - 100},
				{at: pairat, bt: pairbt - 10000},
				{at: pairat, bt: pairbt - 1000000},

				{at: pairat + 1, bt: pairbt},
				{at: pairat + 100, bt: pairbt},
				{at: pairat + 10000, bt: pairbt},
				{at: pairat + 1000000, bt: pairbt},

				{at: pairat - 1, bt: pairbt},
				{at: pairat - 100, bt: pairbt},
				{at: pairat - 10000, bt: pairbt},
				{at: pairat - 1000000, bt: pairbt},

				{at: pairat - 1, bt: pairbt - 1},
				{at: pairat - 100, bt: pairbt - 100},
				{at: pairat - 10000, bt: pairbt - 10000},
				{at: pairat - 1000000, bt: pairbt - 1000000},

				{at: pairat + 1, bt: pairbt + 1},
				{at: pairat + 100, bt: pairbt + 100},
				{at: pairat + 10000, bt: pairbt + 10000},
				{at: pairat + 1000000, bt: pairbt + 1000000},

				{at: pairat - 1, bt: pairbt + 1},
				{at: pairat - 100, bt: pairbt + 100},
				{at: pairat - 10000, bt: pairbt + 10000},
				{at: pairat - 1000000, bt: pairbt + 1000000},

				{at: pairat + 1, bt: pairbt - 1},
				{at: pairat + 100, bt: pairbt - 100},
				{at: pairat + 10000, bt: pairbt - 10000},
				{at: pairat + 1000000, bt: pairbt - 1000000},
			} {
				at := float64(tr.at)
				bt := float64(tr.bt)
				if at < 0 || bt <= at {
					continue
				}

				apx := float64(cd.a.pos.X) + at*float64(cd.a.vel.X)
				apy := float64(cd.a.pos.Y) + at*float64(cd.a.vel.Y)
				apz := float64(cd.a.pos.Z) + at*float64(cd.a.vel.Z)
				bpx := float64(cd.b.pos.X) + bt*float64(cd.b.vel.X)
				bpy := float64(cd.b.pos.Y) + bt*float64(cd.b.vel.Y)
				bpz := float64(cd.b.pos.Z) + bt*float64(cd.b.vel.Z)

				dt := bt - at
				dpx := bpx - apx
				dpy := bpy - apy
				dpz := bpz - apz

				vx := dpx / dt
				vy := dpy / dt
				vz := dpz / dt
				px := bpx - vx*bt
				py := bpy - vy*bt
				pz := bpz - vz*bt

				h := Hailstone{
					pos: space.Point3{
						X: int(math.Round(px)),
						Y: int(math.Round(py)),
						Z: int(math.Round(pz)),
					},
					vel: space.Point3{
						X: int(math.Round(vx)),
						Y: int(math.Round(vy)),
						Z: int(math.Round(vz)),
					},
				}
				loss := totalLoss(h)
				if loss < tryMinLoss {
					tryMinLoss = loss
					tryMin = h
					tryat = tr.at
					trybt = tr.bt
				}
			}

			if tryMinLoss < pairMinLoss {
				pairMinLoss = tryMinLoss
				pairMin = tryMin
				pairat = tryat
				pairbt = trybt
			} else {
				break
			}
		}

		if pairMinLoss < minLoss {
			minLoss = pairMinLoss
			min = pairMin

			if minLoss < 0.5 {
				break
			}
		}
	}
	log.Println(min, minLoss)
	log.Part2(min.pos.X + min.pos.Y + min.pos.Z)
}
