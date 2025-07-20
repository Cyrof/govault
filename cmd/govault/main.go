package main

import (
	"fmt"
	"github.com/Cyrof/govault/internal/fileIO"
)

func main() {
	io := fileIO.NewFileIO()
	if io.CheckMetaFile() {
		fmt.Println("Meta file exists")
		data, err := io.ReadMeta()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(data))
		}
	} else {
		io.EnsureVaultDir()
		fmt.Println("Vault directory created!!")

		metaData := []byte(`{"salt":"abc123","hash":"xyz789"}`)
		err := io.WriteMeta(metaData)
		if err != nil {
			fmt.Println("Failed to write meta:", err)
		}
	}
}
