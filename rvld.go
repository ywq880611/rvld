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

	objfile := linker.NewObjectFile(file)
	//utils.Assert(len(inputfile.ElfSections) == 11, "wrong section headers.")
	objfile.Parse()

	for _, shdr := range objfile.ElfSections {
		println(linker.ElfGetName(objfile.ShStrTable, shdr.Name))
	}

	println(objfile.FirstGlobal)
	println(len(objfile.ElfSyms))

	for _, sym := range objfile.ElfSyms {
		println(linker.ElfGetName(objfile.SymStrTable, sym.Name))
	}

}
