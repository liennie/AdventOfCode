package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type step struct {
	min, max util.Point3
	on       bool
}

func parse(filename string) []step {
	res := []step{}
	for line := range load.File(filename) {
		s := step{}

		if strings.HasPrefix(line, "on") {
			s.on = true
			line = strings.TrimPrefix(line, "on ")
		} else {
			line = strings.TrimPrefix(line, "off ")
		}

		for _, c := range strings.SplitN(line, ",", 3) {
			p := strings.SplitN(c, "=", 2)
			m := util.SplitN(p[1], "..", 2)

			switch p[0] {
			case "x":
				s.min.X = util.Min(m[0], m[1])
				s.max.X = util.Max(m[0], m[1])
			case "y":
				s.min.Y = util.Min(m[0], m[1])
				s.max.Y = util.Max(m[0], m[1])
			case "z":
				s.min.Z = util.Min(m[0], m[1])
				s.max.Z = util.Max(m[0], m[1])
			}
		}

		res = append(res, s)
	}
	return res
}

func reboot(steps []step, min, max util.Point3) map[util.Point3]bool {
	res := map[util.Point3]bool{}

	for _, step := range steps {
		for x := util.Max(step.min.X, min.X); x <= util.Min(step.max.X, max.X); x++ {
			for y := util.Max(step.min.Y, min.Y); y <= util.Min(step.max.Y, max.Y); y++ {
				for z := util.Max(step.min.Z, min.Z); z <= util.Min(step.max.Z, max.Z); z++ {
					if step.on {
						res[util.Point3{X: x, Y: y, Z: z}] = true
					} else {
						delete(res, util.Point3{X: x, Y: y, Z: z})
					}
				}
			}
		}
	}

	return res
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	steps := parse(filename)

	// Part 1
	log.Part1(len(reboot(steps, util.Point3{X: -50, Y: -50, Z: -50}, util.Point3{X: 50, Y: 50, Z: 50})))
}
