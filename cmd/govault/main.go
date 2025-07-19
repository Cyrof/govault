package main

import (
	"github.com/Cyrof/govault/internal/fileIO"
)

func main() {
	io := fileIO.NewFileIO()
	io.PrintPaths()
	/* if io.CheckMetaFile() {
		fmt.Println("Meta file exists")
	} else {
		io.EnsureVaultDir()
		fmt.Println("Vault directory created!!")
		io.PrintPaths()
	} */
}
