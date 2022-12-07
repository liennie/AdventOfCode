package main

import (
	"fmt"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type number interface {
	parent() number
	setParent(number)
	level() int
	magnitude() int
	clone() number
}

type node struct {
	par number
}

func (n *node) parent() number {
	return n.par
}

func (n *node) setParent(parent number) {
	n.par = parent
}

func (n *node) level() int {
	if n.par == nil {
		return 0
	}
	return n.par.level() + 1
}

type rootNode struct {
	child number
}

var _ number = &rootNode{}

func newRoot(child number) *rootNode {
	res := &rootNode{
		child: child,
	}

	res.child.setParent(res)

	return res
}

func (n *rootNode) String() string {
	return fmt.Sprint(n.child)
}

func (n *rootNode) parent() number {
	return nil
}

func (n *rootNode) setParent(parent number) {
	evil.Panic("Cannot set parent on root")
}

func (n *rootNode) level() int {
	return -1
}

func (n *rootNode) magnitude() int {
	return n.child.magnitude()
}

func (n *rootNode) clone() number {
	return newRoot(n.child.clone())
}

type pair struct {
	node
	left, right number
}

var _ number = &pair{}

func newPair(left, right number) *pair {
	res := &pair{
		left:  left,
		right: right,
	}

	res.left.setParent(res)
	res.right.setParent(res)

	return res
}

func (p *pair) String() string {
	return fmt.Sprintf("[%s,%s]", p.left, p.right)
}

func (p *pair) magnitude() int {
	return 3*p.left.magnitude() + 2*p.right.magnitude()
}

func (p *pair) clone() number {
	return newPair(p.left.clone(), p.right.clone())
}

type regular struct {
	node
	val int
}

var _ number = &regular{}

func newRegular(val int) *regular {
	return &regular{
		val: val,
	}
}

func (r *regular) String() string {
	return fmt.Sprintf("%d", r.val)
}

func (r *regular) magnitude() int {
	return r.val
}

func (r *regular) clone() number {
	return newRegular(r.val)
}

func expect(num string, s string) string {
	if strings.HasPrefix(num, s) {
		return num[len(s):]
	}
	evil.Panic("Expected %s", s)
	return ""
}

func startsWith(num string, s string) (bool, string) {
	if strings.HasPrefix(num, s) {
		return true, num[len(s):]
	}
	return false, num
}

func parseNumber(num string) (number, string) {
	if ok, num := startsWith(num, "["); ok {
		left, num := parseNumber(num)
		num = expect(num, ",")
		right, num := parseNumber(num)
		num = expect(num, "]")

		return newPair(left, right), num
	} else {
		return newRegular(int(num[0] - '0')), num[1:]
	}
}

func parse(filename string) []*rootNode {
	res := []*rootNode{}
	for line := range load.File(filename) {
		num, r := parseNumber(line)
		if r != "" {
			evil.Panic("Leftovers: %q", r)
		}
		res = append(res, newRoot(num))
	}
	return res
}

func leftInnerRegular(from number) *regular {
	switch from := from.(type) {
	case *rootNode:
		return leftInnerRegular(from.child)
	case *pair:
		return leftInnerRegular(from.left)
	case *regular:
		return from
	}

	return nil
}

func rightInnerRegular(from number) *regular {
	switch from := from.(type) {
	case *rootNode:
		return rightInnerRegular(from.child)
	case *pair:
		return rightInnerRegular(from.right)
	case *regular:
		return from
	}

	return nil
}

func leftParent(from number) *pair {
	if parent, ok := from.parent().(*pair); ok {
		if parent.left == from {
			return parent
		}
		return leftParent(parent)
	}
	return nil
}

func rightParent(from number) *pair {
	if parent, ok := from.parent().(*pair); ok {
		if parent.right == from {
			return parent
		}
		return rightParent(parent)
	}
	return nil
}

func leftRegular(from number) *regular {
	if p := rightParent(from); p != nil {
		return rightInnerRegular(p.left)
	}
	return nil
}

func rightRegular(from number) *regular {
	if p := leftParent(from); p != nil {
		return leftInnerRegular(p.right)
	}
	return nil
}

func explode(n *pair) {
	if l := leftRegular(n); l != nil {
		l.val += n.left.magnitude()
	}
	if r := rightRegular(n); r != nil {
		r.val += n.right.magnitude()
	}

	replace(n, newRegular(0))
}

func split(n *regular) {
	replace(n, newPair(
		newRegular(n.val/2),
		newRegular((n.val+1)/2),
	))
}

func replace(from, to number) {
	if from.parent() == nil {
		evil.Panic("Cannot replace root node")
	}

	switch parent := from.parent().(type) {
	case *rootNode:
		parent.child = to
		to.setParent(parent)

	case *pair:
		if parent.left == from {
			parent.left = to
		} else if parent.right == from {
			parent.right = to
		}
		to.setParent(parent)
	}
}

func explodeLeftmost(n number) bool {
	switch n := n.(type) {
	case *rootNode:
		return explodeLeftmost(n.child)

	case *pair:
		if n.level() == 4 {
			explode(n)
			return true
		}

		if explodeLeftmost(n.left) {
			return true
		}
		if explodeLeftmost(n.right) {
			return true
		}
	}

	return false
}

func splitLeftmost(n number) bool {
	switch n := n.(type) {
	case *rootNode:
		return splitLeftmost(n.child)

	case *pair:
		if splitLeftmost(n.left) {
			return true
		}
		if splitLeftmost(n.right) {
			return true
		}

	case *regular:
		if n.val >= 10 {
			split(n)
			return true
		}
	}

	return false
}

func reduce(n number) {
	ok := true
	for ok {
		if ok = explodeLeftmost(n); !ok {
			ok = splitLeftmost(n)
		}
	}
}

func add(a, b *rootNode) *rootNode {
	res := newRoot(newPair(
		a.child.clone(),
		b.child.clone(),
	))
	reduce(res)
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	numbers := parse(filename)

	// Part 1
	res := numbers[0]
	for i := 1; i < len(numbers); i++ {
		res = add(res, numbers[i])
	}
	log.Part1(res.magnitude())

	// Part 2
	max := 0
	for i := range numbers {
		for j := range numbers {
			if i == j {
				continue
			}

			if mag := add(numbers[i], numbers[j]).magnitude(); mag > max {
				max = mag
			}
			if mag := add(numbers[j], numbers[i]).magnitude(); mag > max {
				max = mag
			}
		}
	}
	log.Part2(max)
}
