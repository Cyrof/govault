package cli

import (
	"bufio"
	"os"
)

var Reader = bufio.NewReader(os.Stdin)
