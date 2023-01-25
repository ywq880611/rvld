package linker

import (
	"github.com/ywq880611/rvld/pkg/utils"
)

type InputFile struct {
	File        *File
	ElfSections []Shdr
}

func NewInputFile(file *File) InputFile {
	f := InputFile{File: file}
	if len(file.Contents) < EhdrSize {
		utils.Fatal("file too small!")
	}
	if !CheckMagic(file.Contents) {
		utils.Fatal("Not an ELF file!")
	}

	ehdr := utils.Read[Ehdr](file.Contents)
	contents := file.Contents[ehdr.ShOff:]
	shdr := utils.Read[Shdr](contents)

	numSections := uint64(ehdr.ShNum)
	if numSections == 0 {
		numSections = shdr.Size
	}

	f.ElfSections = []Shdr{shdr}

	for numSections > 1 {
		contents = contents[ShdrSize:]
		f.ElfSections = append(f.ElfSections, utils.Read[Shdr](contents))
		numSections--
	}

	return f
}
