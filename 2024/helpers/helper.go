package helpers

import (
	"fmt"
	"time"
)

type ColorStruct struct {
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	White   string
	Gray    string
}

var Color = ColorStruct{
	Red:     "\033[31m",
	Green:   "\033[32m",
	Yellow:  "\033[33m",
	Blue:    "\033[34m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	White:   "\033[37m",
	Gray:    "\033[90m",
}

const Reset = "\033[0m"

func ClearScreen() {
	fmt.Print("\033[2J")
}

func MoveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func Sleep() {
	time.Sleep(100 * time.Millisecond)
}
