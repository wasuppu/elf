package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

type ElfParser struct {
	file        *os.File
	ehdr        *Elf64Header
	phdrs       []*Elf64ProgramHeader
	shdrDesps   []*Elf64SectionHeaderDesp
	symbolDesps []*Elf64SymbolHeaderDesp
	order       binary.ByteOrder
}

func (p ElfParser) PrintEhdr() {
	fmt.Println(p.GetEhdr())
}

func (p *ElfParser) GetEhdr() *Elf64Header {
	if p.ehdr != nil {
		return p.ehdr
	}
	header := new(Elf64Header)
	read(p.order, p.file, 0, header)
	p.ehdr = header
	return header
}

func (p ElfParser) PrintPhdrs() {
	fmt.Printf("\nElf file type is %s\n", e_type[p.ehdr.E_type])
	fmt.Printf("Entry point 0x%x\n", p.ehdr.E_entry)
	phdrs := p.GetPhdrs()
	fmt.Printf("There are %d program headers, starting at offset %v\n\n", len(phdrs), p.ehdr.E_phoff)

	fmt.Println("Program Headers:")
	fmt.Println("  Type           Offset             VirtAddr           PhysAddr")
	fmt.Println("                 FileSiz            MemSiz              Flags  Align")
	for _, phdr := range phdrs {
		fmt.Println(phdr)
	}
}

func (p *ElfParser) GetPhdrs() []*Elf64ProgramHeader {
	if len(p.phdrs) != 0 {
		return p.phdrs
	}

	for i := int64(0); i < int64(p.ehdr.E_phnum); i++ {
		phdr := new(Elf64ProgramHeader)
		pos := int64(p.ehdr.E_phoff) + int64(Elf64_Off(p.ehdr.E_phentsize))*i
		read(p.order, p.file, pos, phdr)
		p.phdrs = append(p.phdrs, phdr)
	}

	return p.phdrs
}

func (p ElfParser) PrintShdrs() {
	fmt.Printf("\nThere are %d section headers, starting at offset 0x%x:\n\n", p.ehdr.E_shnum, p.ehdr.E_shoff)

	fmt.Println("Section Headers:")
	fmt.Println("  [Nr] Name              Type             Address           Offset")
	fmt.Println("       Size              EntSize          Flags  Link  Info  Align")

	desps := p.GetShdrs()
	for _, desp := range desps {
		fmt.Println(desp)
	}
}

func (p *ElfParser) GetShdrs() []*Elf64SectionHeaderDesp {
	if len(p.shdrDesps) != 0 {
		return p.shdrDesps
	}

	shnum := int64(p.ehdr.E_shnum)
	shoff := int64(p.ehdr.E_shoff)
	shentsize := int64(p.ehdr.E_shentsize)
	shstrndx := int64(p.ehdr.E_shstrndx)

	var shstrtaboffset int64
	offset := shoff + shentsize*shstrndx

	if shnum != 0 && shstrndx < shnum {
		shdr := new(Elf64SectionHeader)
		read(p.order, p.file, offset, shdr)
		shstrtaboffset = int64(shdr.SH_offset)
	}

	for i := range shnum {
		shdr := new(Elf64SectionHeader)
		desp := new(Elf64SectionHeaderDesp)

		pos := shoff + shentsize*i
		read(p.order, p.file, pos, shdr)

		desp.shdr = shdr
		desp.idx = int(i)

		// get string of sh_name
		_, err := p.file.Seek(int64(shdr.SH_name)+shstrtaboffset, io.SeekStart)
		check(err)
		reader := bufio.NewReader(p.file)
		desp.name, _ = reader.ReadString(0)

		p.shdrDesps = append(p.shdrDesps, desp)
	}

	return p.shdrDesps
}

func (p ElfParser) PrintSyms() {
	desps := p.GetSyms()
	fmt.Printf("\nSymbol table '.symtab' contains %d entries:\n", len(desps))
	fmt.Println("   Num:    Value          Size Type    Bind   Vis      Ndx Name")
	for _, desp := range desps {
		fmt.Println(desp)
	}
}

func (p *ElfParser) GetSyms() []*Elf64SymbolHeaderDesp {
	if len(p.symbolDesps) != 0 {
		return p.symbolDesps
	}

	if len(p.shdrDesps) == 0 {
		p.GetShdrs()
	}

	var strtaboffset int64
	var dynstroffset int64
	for _, shdrDesp := range p.shdrDesps {
		shdr := shdrDesp.shdr
		if shdr.SH_type == SHT_STRTAB && strings.Trim(shdrDesp.name, "\x00") == ".strtab" {
			strtaboffset = int64(shdr.SH_offset)
			break
		}
	}

	for _, shdrDesp := range p.shdrDesps {
		shdr := shdrDesp.shdr
		if shdr.SH_type == SHT_STRTAB && strings.Trim(shdrDesp.name, "\x00") == ".dynstr" {
			dynstroffset = int64(shdr.SH_offset)
			break
		}
	}

	for _, shdrDesp := range p.shdrDesps {
		shdr := shdrDesp.shdr
		if shdr.SH_type != SHT_SYMTAB && shdr.SH_type != SHT_DYNSYM {
			continue
		}

		/*
			// The code in the comments gets the same result of strtaboffset as the code in the loop at the beginning of this function, just calculated it again
			strtabidx := int64(shdr.SH_link)
			strtabshdroffset := int64(p.ehdr.E_shoff) + int64(p.ehdr.E_shentsize)*strtabidx
			_, err := p.file.Seek(strtabshdroffset, io.SeekStart)
			check(err)
			tabshdr := new(Elf64SectionHeader)
			read[Elf64SectionHeader](p.file, strtabshdroffset, tabshdr)
			strtaboffset = int64(tabshdr.SH_offset)
			fmt.Println(strtaboffset)
		*/

		symNum := int(shdr.SH_size) / int(shdr.SH_entsize)
		for i := range symNum {
			symbol := new(Elf64SymbolHeader)
			desp := new(Elf64SymbolHeaderDesp)

			symoffset := int64(shdr.SH_offset) + int64(i)*int64(shdr.SH_entsize)
			read(p.order, p.file, symoffset, symbol)

			symnameoffset := int64(symbol.ST_name)
			if shdr.SH_type == SHT_SYMTAB {
				_, err := p.file.Seek(strtaboffset+symnameoffset, io.SeekStart)
				check(err)
			}

			if shdr.SH_type == SHT_DYNSYM {
				_, err := p.file.Seek(dynstroffset+symnameoffset, io.SeekStart)
				check(err)
			}
			reader := bufio.NewReader(p.file)
			desp.name, _ = reader.ReadString(0)

			desp.idx = i
			desp.sym = symbol

			p.symbolDesps = append(p.symbolDesps, desp)
		}

	}

	return p.symbolDesps
}

func LoadData(file *os.File) (*ElfParser, error) {
	p := new(ElfParser)
	p.file = file

	if err := p.testElf(); err != nil {
		return nil, err
	}

	p.GetEhdr()

	return p, nil
}

func (p *ElfParser) testElf() error {
	_, err := p.file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	ident := make([]byte, EI_NIDENT)
	_, err = p.file.Read(ident)
	if err != nil {
		return err
	}

	if string(ident[0:SELFMAG]) != ELFMAG {
		return fmt.Errorf("error: not an elf file - it has the wrong magic bytes at the start")
	}

	if ident[EI_CLASS] != ELFCLASS64 {
		return fmt.Errorf("only support 64-bit files")
	}

	if ei_data[Elf_UChar(ident[EI_DATA])] == "LSB" {
		p.order = binary.LittleEndian
	} else if ei_data[Elf_UChar(ident[EI_DATA])] == "MSB" {
		p.order = binary.BigEndian
	} else {
		p.order = binary.LittleEndian
	}

	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read[T any](order binary.ByteOrder, file *os.File, pos int64, obj *T) *T {
	_, err := file.Seek(pos, io.SeekStart)
	check(err)
	err = binary.Read(file, order, obj)
	check(err)
	return obj
}

/* only work for .strtab, .shstrtab */
/*
func (p *ElfParser) dump(name string) []string {
	if len(p.shdrDesps) == 0 {
		p.GetShdr()
	}
	strtab := []string{}

	for _, shdrDesp := range p.shdrDesps {
		shdr := shdrDesp.shdr
		if strings.Trim(shdrDesp.name, "\x00") == name {
			offset := int64(shdr.SH_offset)
			size := int(shdr.SH_size)

			_, err := p.file.Seek(offset, io.SeekStart)
			check(err)
			reader := bufio.NewReader(p.file)
			l := 0
			for l <= size {
				s, err := reader.ReadString(0)
				check(err)

				strtab = append(strtab, s)
				l += len(s)
			}
			break
		}
	}
	return strtab
}
*/
