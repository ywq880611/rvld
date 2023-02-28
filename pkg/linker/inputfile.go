package linker

import (
	"debug/elf"
	"fmt"

	"github.com/ywq880611/rvld/pkg/utils"
)

type InputFile struct {
	File        *File
	ElfSections []Shdr
	ElfSyms     []Sym
	FirstGlobal int64
	ShStrTable  []byte
	SymStrTable []byte
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

	shstrndx := uint32(ehdr.ShStrndx)
	if shstrndx == uint32(elf.SHN_XINDEX) {
		shstrndx = uint32(shdr.Link)
	}

	//println(shstrndx)
	f.ShStrTable = f.GetBytesFromIdx(shstrndx)

	//fmt.Printf("%s\n", f.ShStrTable)
	return f
}

func (f *InputFile) GetBytesFromShdr(s *Shdr) []byte {
	end := s.Offset + s.Size
	if uint64(len(f.File.Contents)) < end {
		utils.Fatal(fmt.Sprintf("section header is out of range , offset is %d", s.Offset))
	}
	return f.File.Contents[s.Offset:end]
}

func (f *InputFile) GetBytesFromIdx(idx uint32) []byte {
	if uint32(len(f.ElfSections)) < idx {
		utils.Fatal(fmt.Sprintf("section idx is out of range , idx is %d", idx))
	}
	return f.GetBytesFromShdr(&f.ElfSections[idx])
}

func (f *InputFile) FillUpElfSyms(s *Shdr) {
	bs := f.GetBytesFromShdr(s)
	sym_num := len(bs) / SymSize
	f.ElfSyms = make([]Sym, 0, sym_num)
	for sym_num > 0 {
		f.ElfSyms = append(f.ElfSyms, utils.Read[Sym](bs))
		bs = bs[SymSize:]
		sym_num--
	}
}

func (f *InputFile) FindSection(ty uint32) *Shdr {
	for i := 0; i < len(f.ElfSections); i++ {
		shdr := &f.ElfSections[i]
		if shdr.Type == ty {
			return shdr
		}
	}
	return nil
}
