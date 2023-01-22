package main

import (
	"os"

	"github.com/ywq880611/rvld/pkg/linker"
	"github.com/ywq880611/rvld/pkg/utils"
)

func main() {
	//println("Hello, world.")
	if len(os.Args) < 2 {
		utils.Fatal("Wrong args")
	}
	//fmt.Printf("%v\n", os.Args)
	file := linker.MustNewFile(os.Args[1])

	if !linker.CheckMagic(file.Contents) {
		utils.Fatal("Not a ELF file")
	}
}
