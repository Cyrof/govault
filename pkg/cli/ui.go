package cli

import (
	"fmt"

	"github.com/fatih/color"
)

func Success(msg string, args ...any) {
	color.Green(fmt.Sprintf(msg, args...))
}

func Error(msg string, args ...any) {
	color.Red(fmt.Sprintf(msg, args...))
}

func Warn(msg string, args ...any) {
	color.Yellow(fmt.Sprintf(msg, args...))
}
