package main

import (
	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func transformStep(init, subject, div int) int {
	value := init
	value *= subject
	value %= div
	return value
}

func transform(subject, div, loop int) int {
	value := 1
	for range loop {
		value = transformStep(value, subject, div)
	}
	return value
}

func parse(filename string) (int, int) {
	ch := load.File(filename)
	defer channel.Drain(ch)

	return evil.Atoi(<-ch), evil.Atoi(<-ch)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	cardKey, doorKey := parse(filename)
	log.Printf("cardKey: %d", cardKey)
	log.Printf("doorKey: %d", doorKey)

	// Part 1
	const subject = 7
	const div = 20201227

	cardLoop, doorLoop := 0, 0

	value := transformStep(1, subject, div)
	loop := 1
	for value != 1 && (cardLoop == 0 || doorLoop == 0) {
		if value == cardKey {
			cardLoop = loop
		}
		if value == doorKey {
			doorLoop = loop
		}
		value = transformStep(value, subject, div)
		loop++
	}

	evil.Assert(cardLoop != 0, "did not find card loop size")
	evil.Assert(doorLoop != 0, "did not find door loop size")

	log.Printf("cardLoop: %d", cardLoop)
	log.Printf("doorLoop: %d", doorLoop)

	log.Part1(transform(cardKey, div, doorLoop))
}
