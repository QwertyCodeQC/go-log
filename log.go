package golog

import (
	"errors"
	"fmt"
	"strings"
)

const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const reset = "\033[0m"

func format(color string, msgType string) string {
	return fmt.Sprintf("%s[%s]%s", color, strings.ToUpper(msgType), reset)
}
func Info(msg string) {
	fmt.Printf("%s %s\n", format(green, "info"), msg)
}

func Error(msg string) {
	fmt.Printf("%s %s\n", format(red, "error"), msg)
}

func Debug(msg string) {
	fmt.Printf("%s %s\n", format(blue, "debug"), msg)
}

func Warn(msg string) {
	fmt.Printf("%s %s\n", format(yellow, "warn"), msg)
}

func Fatal(msg string) {
	fmt.Printf("%s %s\n", format(red, "fatal"), msg)
}
func ErrorLadder(err error) {
	if err == nil {
		return
	}

	fmt.Printf("%s%s\n", format(red, "error: "), err) // print main error once

	unwrapped := errors.Unwrap(err)
	i := 1

	for unwrapped != nil {
		prefix := fmt.Sprintf("  %s↳ ", string(make([]byte, i*2)))
		fmt.Printf("%s[%d] %T: %v\n", prefix, i, unwrapped, unwrapped)
		unwrapped = errors.Unwrap(unwrapped)
		i++
	}
}
