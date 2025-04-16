package main

import (
	"nas-file-tool/internal"
	"nas-file-tool/pkg/utils"
	"os"
)

func main() {
	if len(os.Args) > 3 {
		rootDir := os.Args[1]
		url1 := os.Args[2]
		url2 := os.Args[3]
		utils.FindMoves(rootDir, []string{url1, url2})
	} else {
		internal.Menu()
	}
}
