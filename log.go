package golog

import (
	"errors"
	"fmt"
)

const RED = "\033[31m"
const GREEN = "\033[32m"
const YELLOW = "\033[33m"
const BLUE = "\033[34m"
const MAGENTA = "\033[35m"
const CYAN = "\033[36m"
const WHITE = "\033[37m"
const RESET = "\033[0m"
const BOLD = "\033[1m"
const UNDERLINE = "\033[4m"
const REVERSED = "\033[7m"
const BLACK = "\033[30m"
const GRAY = "\033[90m"
func format(color string, msgType string) string {
	return fmt.Sprintf("%s%s:%s", color, msgType, RESET)
}
func Info() {
	fmt.Println(format(GREEN, "info"))
}

func Error() {
	fmt.Println(format(RED, "error"))
}

func Debug() {
	fmt.Println(format(BLUE, "debug"))
}

func Warn() {
	fmt.Println(format(YELLOW, "warn"))
}
func Fatal() {
	fmt.Println(format(RED, "fatal"))
}
func ErrorLadder(err error) {
	if err == nil {
		return
	}

	fmt.Printf("%s%s\n", format(RED, "error: "), err) // print main error once

	unwrapped := errors.Unwrap(err)
	i := 1

	for unwrapped != nil {
		prefix := fmt.Sprintf("  %s↳ ", string(make([]byte, i*2)))
		fmt.Printf("%s[%d] %T: %v\n", prefix, i, unwrapped, unwrapped)
		unwrapped = errors.Unwrap(unwrapped)
		i++
	}
}

