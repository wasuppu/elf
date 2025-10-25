package main

import (
	"bytes"
	"fmt"
)

type Elf64_Addr uint64
type Elf64_Half uint16
type Elf64_Off uint64
type Elf64_Word uint32
type Elf64_SWord int32
type Elf64_XWord uint64
type Elf64_SXWord int64

type Elf_UChar uint8 // Unsigned small integer

// ELF Object File Format
// Linking View:    ELF header | Program header table(optional) | seciton(n) | section header table
// Execution View:  ELF header | Program header table           | segment(n) | section header table (optional)

// Executable Headers (Ehdr)
type Elf64Header struct {
	E_ident     [EI_NIDENT]Elf_UChar /* ELF identification */
	E_type      Elf64_Half           /* Object file type */
	E_machine   Elf64_Half           /* Architecture */
	E_version   Elf64_Word           /* Object file version */
	E_entry     Elf64_Addr           /* Entry point virtual address */
	E_phoff     Elf64_Off            /* Program header offset */
	E_shoff     Elf64_Off            /* Section header offset */
	E_flags     Elf64_Word           /* Processor-specific flags */
	E_ehsize    Elf64_Half           /* ELF header size */
	E_phentsize Elf64_Half           /* Size of program header entry */
	E_phnum     Elf64_Half           /* Number of program header entries */
	E_shentsize Elf64_Half           /* Size of section header entry */
	E_shnum     Elf64_Half           /* Number of section header entries */
	E_shstrndx  Elf64_Half           /* Section name string table index */
}

func (ehdr Elf64Header) String() string {
	hexString := ""
	for _, b := range ehdr.E_ident {
		hexString += fmt.Sprintf("%02x ", b)
	}

	// e_indent
	builder := bytes.NewBuffer([]byte{})
	builder.WriteString("ELF Header:\n")
	fmt.Fprintf(builder, "  %-40s%v\n", "Magic:", hexString)
	fmt.Fprintf(builder, "  %-40s%v\n", "Class:", ei_class[ehdr.E_ident[EI_CLASS]])
	fmt.Fprintf(builder, "  %-40s%v\n", "Data:", ei_data[ehdr.E_ident[EI_DATA]])
	fmt.Fprintf(builder, "  %-40s%v\n", "Version:", ehdr.E_ident[EI_VERSION])
	fmt.Fprintf(builder, "  %-40s%v\n", "OS/ABI:", ei_osabi[ehdr.E_ident[EI_OSABI]])
	fmt.Fprintf(builder, "  %-40s%v\n", "ABI version:", ehdr.E_ident[EI_ABIVERSION])
	fmt.Fprintf(builder, "  %-40s%v\n", "Byte index:", ehdr.E_ident[EI_PAD])

	fmt.Fprintf(builder, "  %-40s%v\n", "Type:", e_type[ehdr.E_type])
	fmt.Fprintf(builder, "  %-40s%v\n", "Machine:", e_machine[ehdr.E_machine])
	fmt.Fprintf(builder, "  %-40s0x%x\n", "Version:", ehdr.E_version)
	fmt.Fprintf(builder, "  %-40s0x%x\n", "Entry point address:", ehdr.E_entry)
	fmt.Fprintf(builder, "  %-40s%v\n", "Program header offset:", ehdr.E_phoff)
	fmt.Fprintf(builder, "  %-40s%v\n", "Section header offset:", ehdr.E_shoff)
	fmt.Fprintf(builder, "  %-40s%v\n", "Flags:", ehdr.E_flags)
	fmt.Fprintf(builder, "  %-40s%v (bytes)\n", "Size of this header", ehdr.E_ehsize)
	fmt.Fprintf(builder, "  %-40s%v (bytes)\n", "Size of program headers", ehdr.E_phentsize)
	fmt.Fprintf(builder, "  %-40s%v\n", "Number of program headers", ehdr.E_phnum)
	fmt.Fprintf(builder, "  %-40s%v (bytes)\n", "Size of section headers", ehdr.E_shentsize)
	fmt.Fprintf(builder, "  %-40s%v\n", "Number of section headers:", ehdr.E_shnum)
	fmt.Fprintf(builder, "  %-40s%v", "Section header string table index:", ehdr.E_shstrndx)
	return builder.String()
}

/* Program segment header (Phdr) */
type Elf64ProgramHeader struct {
	P_type   Elf64_Word  /* Segment type */
	P_flags  Elf64_Word  /* Segment flags */
	P_offset Elf64_Off   /* Segment file offset */
	P_vaddr  Elf64_Addr  /* Segment virtual address */
	P_paddr  Elf64_Addr  /* Segment physical address */
	P_filesz Elf64_XWord /* Segment size in file */
	P_memsz  Elf64_XWord /* Segment size in memory */
	P_align  Elf64_XWord /* Segment alignment */
}

func (phdr Elf64ProgramHeader) String() string {
	builder := bytes.NewBuffer([]byte{})
	fmt.Fprintf(builder, "  %-14v ", p_type[phdr.P_type])
	fmt.Fprintf(builder, "0x%016x ", phdr.P_offset)
	fmt.Fprintf(builder, "0x%016x ", phdr.P_vaddr)
	fmt.Fprintf(builder, "0x%016x\n", phdr.P_paddr)
	fmt.Fprintf(builder, "%17s0x%016x ", "", phdr.P_filesz)
	fmt.Fprintf(builder, "0x%016x ", phdr.P_memsz)
	fmt.Fprintf(builder, " %-9v", getSegmentFlags(phdr.P_flags))
	fmt.Fprintf(builder, "0x%x", phdr.P_align)
	return builder.String()
}

/* Section header (Shdr) */
type Elf64SectionHeader struct {
	SH_name      Elf64_Word  /* Section name */
	SH_type      Elf64_Word  /* Section type */
	SH_flags     Elf64_XWord /* Section attributes */
	SH_addr      Elf64_Addr  /* Virtual address in memory */
	SH_offset    Elf64_Off   /* Offset in file */
	SH_size      Elf64_XWord /* Size of section */
	SH_link      Elf64_Word  /* Link to other section */
	SH_info      Elf64_Word  /* Miscellaneous information */
	SH_addralign Elf64_XWord /* Address alignment boundary */
	SH_entsize   Elf64_XWord /* Size of entries, if section has table */
}

func (shdr Elf64SectionHeader) String() string {
	builder := bytes.NewBuffer([]byte{})
	fmt.Fprintf(builder, "%-016v ", sh_type[shdr.SH_type])
	fmt.Fprintf(builder, "%016x  ", shdr.SH_addr)
	fmt.Fprintf(builder, "%08x\n", shdr.SH_offset)

	fmt.Fprintf(builder, "%7s%016x  ", "", shdr.SH_size)
	fmt.Fprintf(builder, "%016x ", shdr.SH_entsize)
	fmt.Fprintf(builder, "%3s", getSectionFlags(shdr.SH_flags))
	fmt.Fprintf(builder, "%8v", shdr.SH_link)
	fmt.Fprintf(builder, "%6v ", shdr.SH_info)
	fmt.Fprintf(builder, "%5v", shdr.SH_addralign)
	return builder.String()
}

/* additional information used to describe the section header */
type Elf64SectionHeaderDesp struct {
	name string
	idx  int
	shdr *Elf64SectionHeader
}

func (d Elf64SectionHeaderDesp) String() string {
	return fmt.Sprintf("  [%2d] %-19s%s", d.idx, d.name, d.shdr)
}

/* Symbol table entry  */
type Elf64SymbolHeader struct {
	ST_name  Elf64_Word  /* Symbol name (string tbl index) */
	ST_info  Elf_UChar   /* Symbol type and binding */
	ST_other Elf_UChar   /* Symbol visibility */
	ST_shndx Elf64_Half  /* Section index */
	ST_value Elf64_Addr  /* Symbol value */
	ST_size  Elf64_XWord /* Symbol size */
}

type Elf64SymbolHeaderDesp struct {
	name string
	idx  int
	sym  *Elf64SymbolHeader
}

func (desp Elf64SymbolHeaderDesp) String() string {
	return fmt.Sprintf("   %3d: %s%s", desp.idx, desp.sym, desp.name)
}

func (sym Elf64SymbolHeader) String() string {
	typ := sym.ST_info & 0xf
	bind := sym.ST_info >> 4
	vis := sym.ST_other & 0x03
	var idx string
	if v, ok := sym_idx[sym.ST_shndx]; ok {
		idx = v
	} else {
		idx = fmt.Sprint(sym.ST_shndx)
	}

	builder := bytes.NewBuffer([]byte{})
	fmt.Fprintf(builder, "%016x  ", sym.ST_value)
	fmt.Fprintf(builder, "%4d ", sym.ST_size)
	fmt.Fprintf(builder, "%-8s", sym_type[typ])
	fmt.Fprintf(builder, "%-7s", sym_bind[bind])
	fmt.Fprintf(builder, "%-9s", sym_vis[vis])
	fmt.Fprintf(builder, "%3s ", idx)
	return builder.String()
}
