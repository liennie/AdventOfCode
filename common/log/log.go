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

func Part1(n interface{}) {
	log.Printf("Part 1: %v", n)
}

func Part2(n interface{}) {
	log.Printf("Part 2: %v", n)
}

func Print(v ...interface{}) {
	log.Print(v...)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}
