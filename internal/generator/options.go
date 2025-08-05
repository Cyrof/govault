package generator

type Options struct {
	Length       int
	UseLowercase bool
	UseUppercase bool
	UseNumbers   bool
	UseSymbols   bool
}

var DefaultOptions = Options{
	Length:       16,
	UseLowercase: true,
	UseUppercase: true,
	UseNumbers:   true,
	UseSymbols:   false,
}
