package generator

import (
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

func BuildCharset(opts Options) string {
	var builder strings.Builder

	if opts.UseLowercase {
		builder.WriteString(lowercase)
	}
	if opts.UseUppercase {
		builder.WriteString(uppercase)
	}
	if opts.UseNumbers {
		builder.WriteString(numbers)
	}
	if opts.UseSymbols {
		builder.WriteString(symbols)
	}

	return builder.String()
}
