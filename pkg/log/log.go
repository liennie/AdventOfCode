package log

import (
	"log"
)

func init() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
}

func Err(err error) {
	log.Printf("Error: %s", err)
}

func Part1(n any) {
	log.Printf("Part 1: %v", n)
}

func Part2(n any) {
	log.Printf("Part 2: %v", n)
}

func Print(v ...any) {
	log.Print(v...)
}

func Printf(format string, v ...any) {
	log.Printf(format, v...)
}

func Println(v ...any) {
	log.Println(v...)
}
