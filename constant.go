package main

/* ELF file header, appears at the start of every ELF file.  */
const EI_NIDENT = 16

/*
Fields in the e_ident array.  The EI_* macros are indices into the array.
The macros under each EI_* macro are the values the byte may have.
*/
const (
	EI_MAG0 = 0    /* File identification byte 0 index */
	ELFMAG0 = 0x7f /* Magic number byte 0 */
	EI_MAG1 = 1    /* File identification byte 1 index */
	ELFMAG1 = 'E'  /* Magic number byte 1 */
	EI_MAG2 = 2    /* File identification byte 2 index */
	ELFMAG2 = 'L'  /* Magic number byte 2 */
	EI_MAG3 = 3    /* File identification byte 3 index */
	ELFMAG3 = 'F'  /* Magic number byte 3 */

	/* Conglomeration of the identification bytes, for easy testing as a word.  */
	ELFMAG  = "\177ELF"
	SELFMAG = 4
)

/* File class byte index */
const EI_CLASS = 4
const (
	ELFCLASSNONE = 0 /* Invalid class */
	ELFCLASS32   = 1 /* 32-bit objects */
	ELFCLASS64   = 2 /* 64-bit objects */
	ELFCLASSNUM  = 3
)

var ei_class = map[Elf_UChar]string{
	ELFCLASSNONE: "None",
	ELFCLASS32:   "32 bits",
	ELFCLASS64:   "64 bits",
}

/* Data encoding byte index */
const EI_DATA = 5
const (
	ELFDATANONE = 0 /* Invalid data encoding */
	ELFDATA2LSB = 1 /* 2's complement, little endian */
	ELFDATA2MSB = 2 /* 2's complement, big endian */
	ELFDATANUM  = 3
)

var ei_data = map[Elf_UChar]string{
	ELFDATANONE: "None",
	ELFDATA2LSB: "LSB",
	ELFDATA2MSB: "MSB",
}

/* File version byte index */
const EI_VERSION = 6

/* OS ABI identification */
const EI_OSABI = 7
const (
	ELFOSABI_NONE       = 0            /* UNIX System V ABI */
	ELFOSABI_SYSV       = 0            /* Alias.  */
	ELFOSABI_HPUX       = 1            /* HP-UX */
	ELFOSABI_NETBSD     = 2            /* NetBSD.  */
	ELFOSABI_GNU        = 3            /* Object uses GNU ELF extensions.  */
	ELFOSABI_LINUX      = ELFOSABI_GNU /* Compatibility alias.  */
	ELFOSABI_SOLARIS    = 6            /* Sun Solaris.  */
	ELFOSABI_AIX        = 7            /* IBM AIX.  */
	ELFOSABI_IRIX       = 8            /* SGI Irix.  */
	ELFOSABI_FREEBSD    = 9            /* FreeBSD.  */
	ELFOSABI_TRU64      = 10           /* Compaq TRU64 UNIX.  */
	ELFOSABI_MODESTO    = 11           /* Novell Modesto.  */
	ELFOSABI_OPENBSD    = 12           /* OpenBSD.  */
	ELFOSABI_ARM_AEABI  = 64           /* ARM EABI */
	ELFOSABI_ARM        = 97           /* ARM */
	ELFOSABI_STANDALONE = 255          /* Standalone (embedded) application */
)

var ei_osabi = map[Elf_UChar]string{
	ELFOSABI_NONE:       "UNIX System V ABI",
	ELFOSABI_HPUX:       "HP-UX",
	ELFOSABI_NETBSD:     "NetBSD",
	ELFOSABI_GNU:        "Object uses GNU ELF extensions",
	ELFOSABI_SOLARIS:    "Sun Solaris",
	ELFOSABI_AIX:        "IBM AIX",
	ELFOSABI_IRIX:       "SGI Irix",
	ELFOSABI_FREEBSD:    "FreeBSD",
	ELFOSABI_TRU64:      "Compaq TRU64 UNIX",
	ELFOSABI_MODESTO:    "Novell Modesto",
	ELFOSABI_OPENBSD:    "OpenBSD",
	ELFOSABI_ARM_AEABI:  "ARM EABI",
	ELFOSABI_ARM:        "ARM",
	ELFOSABI_STANDALONE: "Standalone (embedded) application",
}

/* ABI version */
const EI_ABIVERSION = 8

/* Byte index of padding bytes */
const EI_PAD = 9

/* End of e_ident */

/* Legal values for e_type (object file type).  */
const (
	ET_NONE   = 0      /* No file type */
	ET_REL    = 1      /* Relocatable file */
	ET_EXEC   = 2      /* Executable file */
	ET_DYN    = 3      /* Shared object file */
	ET_CORE   = 4      /* Core file */
	ET_NUM    = 5      /* Number of defined types */
	ET_LOOS   = 0xfe00 /* OS-specific range start */
	ET_HIOS   = 0xfeff /* OS-specific range end */
	ET_LOPROC = 0xff00 /* Processor-specific range start */
	ET_HIPROC = 0xffff /* Processor-specific range end */
)

var e_type = map[Elf64_Half]string{
	ET_NONE: "NONE (None)",
	ET_REL:  "REL (Relocatable file)",
	ET_EXEC: "EXEC (Executable file)",
	ET_DYN:  "DYN",
	ET_CORE: "CORE (Core file)",
}

/* Legal values for e_machine (architecture).  */
const (
	EM_NONE        = 0  /* No machine */
	EM_M32         = 1  /* AT&T WE 32100 */
	EM_SPARC       = 2  /* SUN SPARC */
	EM_386         = 3  /* Intel 80386 */
	EM_68K         = 4  /* Motorola m68k family */
	EM_88K         = 5  /* Motorola m88k family */
	EM_IAMCU       = 6  /* Intel MCU */
	EM_860         = 7  /* Intel 80860 */
	EM_MIPS        = 8  /* MIPS R3000 big-endian */
	EM_S370        = 9  /* IBM System/370 */
	EM_MIPS_RS3_LE = 10 /* MIPS R3000 little-endian */
	/* reserved 11-14 */
	EM_PARISC = 15 /* HPPA */
	/* reserved 16 */
	EM_VPP500      = 17 /* Fujitsu VPP500 */
	EM_SPARC32PLUS = 18 /* Sun's "v8plus" */
	EM_960         = 19 /* Intel 80960 */
	EM_PPC         = 20 /* PowerPC */
	EM_PPC64       = 21 /* PowerPC 64-bit */
	EM_S390        = 22 /* IBM S390 */
	EM_SPU         = 23 /* IBM SPU/SPC */
	/* reserved 24-35 */
	EM_V800         = 36  /* NEC V800 series */
	EM_FR20         = 37  /* Fujitsu FR20 */
	EM_RH32         = 38  /* TRW RH-32 */
	EM_RCE          = 39  /* Motorola RCE */
	EM_ARM          = 40  /* ARM */
	EM_FAKE_ALPHA   = 41  /* Digital Alpha */
	EM_SH           = 42  /* Hitachi SH */
	EM_SPARCV9      = 43  /* SPARC v9 64-bit */
	EM_TRICORE      = 44  /* Siemens Tricore */
	EM_ARC          = 45  /* Argonaut RISC Core */
	EM_H8_300       = 46  /* Hitachi H8/300 */
	EM_H8_300H      = 47  /* Hitachi H8/300H */
	EM_H8S          = 48  /* Hitachi H8S */
	EM_H8_500       = 49  /* Hitachi H8/500 */
	EM_IA_64        = 50  /* Intel Merced */
	EM_MIPS_X       = 51  /* Stanford MIPS-X */
	EM_COLDFIRE     = 52  /* Motorola Coldfire */
	EM_68HC12       = 53  /* Motorola M68HC12 */
	EM_MMA          = 54  /* Fujitsu MMA Multimedia Accelerator */
	EM_PCP          = 55  /* Siemens PCP */
	EM_NCPU         = 56  /* Sony nCPU embeeded RISC */
	EM_NDR1         = 57  /* Denso NDR1 microprocessor */
	EM_STARCORE     = 58  /* Motorola Start*Core processor */
	EM_ME16         = 59  /* Toyota ME16 processor */
	EM_ST100        = 60  /* STMicroelectronic ST100 processor */
	EM_TINYJ        = 61  /* Advanced Logic Corp. Tinyj emb.fam */
	EM_X86_64       = 62  /* AMD x86-64 architecture */
	EM_PDSP         = 63  /* Sony DSP Processor */
	EM_PDP10        = 64  /* Digital PDP-10 */
	EM_PDP11        = 65  /* Digital PDP-11 */
	EM_FX66         = 66  /* Siemens FX66 microcontroller */
	EM_ST9PLUS      = 67  /* STMicroelectronics ST9+ 8/16 mc */
	EM_ST7          = 68  /* STmicroelectronics ST7 8 bit mc */
	EM_68HC16       = 69  /* Motorola MC68HC16 microcontroller */
	EM_68HC11       = 70  /* Motorola MC68HC11 microcontroller */
	EM_68HC08       = 71  /* Motorola MC68HC08 microcontroller */
	EM_68HC05       = 72  /* Motorola MC68HC05 microcontroller */
	EM_SVX          = 73  /* Silicon Graphics SVx */
	EM_ST19         = 74  /* STMicroelectronics ST19 8 bit mc */
	EM_VAX          = 75  /* Digital VAX */
	EM_CRIS         = 76  /* Axis Communications 32-bit emb.proc */
	EM_JAVELIN      = 77  /* Infineon Technologies 32-bit emb.proc */
	EM_FIREPATH     = 78  /* Element 14 64-bit DSP Processor */
	EM_ZSP          = 79  /* LSI Logic 16-bit DSP Processor */
	EM_MMIX         = 80  /* Donald Knuth's educational 64-bit proc */
	EM_HUANY        = 81  /* Harvard University machine-independent object files */
	EM_PRISM        = 82  /* SiTera Prism */
	EM_AVR          = 83  /* Atmel AVR 8-bit microcontroller */
	EM_FR30         = 84  /* Fujitsu FR30 */
	EM_D10V         = 85  /* Mitsubishi D10V */
	EM_D30V         = 86  /* Mitsubishi D30V */
	EM_V850         = 87  /* NEC v850 */
	EM_M32R         = 88  /* Mitsubishi M32R */
	EM_MN10300      = 89  /* Matsushita MN10300 */
	EM_MN10200      = 90  /* Matsushita MN10200 */
	EM_PJ           = 91  /* picoJava */
	EM_OPENRISC     = 92  /* OpenRISC 32-bit embedded processor */
	EM_ARC_COMPACT  = 93  /* ARC International ARCompact */
	EM_XTENSA       = 94  /* Tensilica Xtensa Architecture */
	EM_VIDEOCORE    = 95  /* Alphamosaic VideoCore */
	EM_TMM_GPP      = 96  /* Thompson Multimedia General Purpose Proc */
	EM_NS32K        = 97  /* National Semi. 32000 */
	EM_TPC          = 98  /* Tenor Network TPC */
	EM_SNP1K        = 99  /* Trebia SNP 1000 */
	EM_ST200        = 100 /* STMicroelectronics ST200 */
	EM_IP2K         = 101 /* Ubicom IP2xxx */
	EM_MAX          = 102 /* MAX processor */
	EM_CR           = 103 /* National Semi. CompactRISC */
	EM_F2MC16       = 104 /* Fujitsu F2MC16 */
	EM_MSP430       = 105 /* Texas Instruments msp430 */
	EM_BLACKFIN     = 106 /* Analog Devices Blackfin DSP */
	EM_SE_C33       = 107 /* Seiko Epson S1C33 family */
	EM_SEP          = 108 /* Sharp embedded microprocessor */
	EM_ARCA         = 109 /* Arca RISC */
	EM_UNICORE      = 110 /* PKU-Unity & MPRC Peking Uni. mc series */
	EM_EXCESS       = 111 /* eXcess configurable cpu */
	EM_DXP          = 112 /* Icera Semi. Deep Execution Processor */
	EM_ALTERA_NIOS2 = 113 /* Altera Nios II */
	EM_CRX          = 114 /* National Semi. CompactRISC CRX */
	EM_XGATE        = 115 /* Motorola XGATE */
	EM_C166         = 116 /* Infineon C16x/XC16x */
	EM_M16C         = 117 /* Renesas M16C */
	EM_DSPIC30F     = 118 /* Microchip Technology dsPIC30F */
	EM_CE           = 119 /* Freescale Communication Engine RISC */
	EM_M32C         = 120 /* Renesas M32C */
	/* r=eserved 121-130 */
	EM_TSK3000       = 131 /* Altium TSK3000 */
	EM_RS08          = 132 /* Freescale RS08 */
	EM_SHARC         = 133 /* Analog Devices SHARC family */
	EM_ECOG2         = 134 /* Cyan Technology eCOG2 */
	EM_SCORE7        = 135 /* Sunplus S+core7 RISC */
	EM_DSP24         = 136 /* New Japan Radio (NJR) 24-bit DSP */
	EM_VIDEOCORE3    = 137 /* Broadcom VideoCore III */
	EM_LATTICEMICO32 = 138 /* RISC for Lattice FPGA */
	EM_SE_C17        = 139 /* Seiko Epson C17 */
	EM_TI_C6000      = 140 /* Texas Instruments TMS320C6000 DSP */
	EM_TI_C2000      = 141 /* Texas Instruments TMS320C2000 DSP */
	EM_TI_C5500      = 142 /* Texas Instruments TMS320C55x DSP */
	EM_TI_ARP32      = 143 /* Texas Instruments App. Specific RISC */
	EM_TI_PRU        = 144 /* Texas Instruments Prog. Realtime Unit */
	/* reserved 145-159 */
	EM_MMDSP_PLUS  = 160 /* STMicroelectronics 64bit VLIW DSP */
	EM_CYPRESS_M8C = 161 /* Cypress M8C */
	EM_R32C        = 162 /* Renesas R32C */
	EM_TRIMEDIA    = 163 /* NXP Semi. TriMedia */
	EM_QDSP6       = 164 /* QUALCOMM DSP6 */
	EM_8051        = 165 /* Intel 8051 and variants */
	EM_STXP7X      = 166 /* STMicroelectronics STxP7x */
	EM_NDS32       = 167 /* Andes Tech. compact code emb. RISC */
	EM_ECOG1X      = 168 /* Cyan Technology eCOG1X */
	EM_MAXQ30      = 169 /* Dallas Semi. MAXQ30 mc */
	EM_XIMO16      = 170 /* New Japan Radio (NJR) 16-bit DSP */
	EM_MANIK       = 171 /* M2000 Reconfigurable RISC */
	EM_CRAYNV2     = 172 /* Cray NV2 vector architecture */
	EM_RX          = 173 /* Renesas RX */
	EM_METAG       = 174 /* Imagination Tech. META */
	EM_MCST_ELBRUS = 175 /* MCST Elbrus */
	EM_ECOG16      = 176 /* Cyan Technology eCOG16 */
	EM_CR16        = 177 /* National Semi. CompactRISC CR16 */
	EM_ETPU        = 178 /* Freescale Extended Time Processing Unit */
	EM_SLE9X       = 179 /* Infineon Tech. SLE9X */
	EM_L10M        = 180 /* Intel L10M */
	EM_K10M        = 181 /* Intel K10M */
	/* reserved 182 */
	EM_AARCH64 = 183 /* ARM AARCH64 */
	/* reserved 184 */
	EM_AVR32       = 185 /* Amtel 32-bit microprocessor */
	EM_STM8        = 186 /* STMicroelectronics STM8 */
	EM_TILE64      = 187 /* Tilera TILE64 */
	EM_TILEPRO     = 188 /* Tilera TILEPro */
	EM_MICROBLAZE  = 189 /* Xilinx MicroBlaze */
	EM_CUDA        = 190 /* NVIDIA CUDA */
	EM_TILEGX      = 191 /* Tilera TILE-Gx */
	EM_CLOUDSHIELD = 192 /* CloudShield */
	EM_COREA_1ST   = 193 /* KIPO-KAIST Core-A 1st gen. */
	EM_COREA_2ND   = 194 /* KIPO-KAIST Core-A 2nd gen. */
	EM_ARCV2       = 195 /* Synopsys ARCv2 ISA.  */
	EM_OPEN8       = 196 /* Open8 RISC */
	EM_RL78        = 197 /* Renesas RL78 */
	EM_VIDEOCORE5  = 198 /* Broadcom VideoCore V */
	EM_78KOR       = 199 /* Renesas 78KOR */
	EM_56800EX     = 200 /* Freescale 56800EX DSC */
	EM_BA1         = 201 /* Beyond BA1 */
	EM_BA2         = 202 /* Beyond BA2 */
	EM_XCORE       = 203 /* XMOS xCORE */
	EM_MCHP_PIC    = 204 /* Microchip 8-bit PIC(r) */
	EM_INTELGT     = 205 /* Intel Graphics Technology */
	/* reserved 206-209 */
	EM_KM32        = 210 /* KM211 KM32 */
	EM_KMX32       = 211 /* KM211 KMX32 */
	EM_EMX16       = 212 /* KM211 KMX16 */
	EM_EMX8        = 213 /* KM211 KMX8 */
	EM_KVARC       = 214 /* KM211 KVARC */
	EM_CDP         = 215 /* Paneve CDP */
	EM_COGE        = 216 /* Cognitive Smart Memory Processor */
	EM_COOL        = 217 /* Bluechip CoolEngine */
	EM_NORC        = 218 /* Nanoradio Optimized RISC */
	EM_CSR_KALIMBA = 219 /* CSR Kalimba */
	EM_Z80         = 220 /* Zilog Z80 */
	EM_VISIUM      = 221 /* Controls and Data Services VISIUMcore */
	EM_FT32        = 222 /* FTDI Chip FT32 */
	EM_MOXIE       = 223 /* Moxie processor */
	EM_AMDGPU      = 224 /* AMD GPU */
	/* reserved 225-242 */
	EM_RISCV = 243 /* RISC-V */

	EM_BPF  = 247 /* Linux BPF -- in-kernel virtual machine */
	EM_CSKY = 252 /* C-SKY */
)

var e_machine = map[Elf64_Half]string{
	EM_NONE:          "No machine",
	EM_M32:           "AT&T WE 32100",
	EM_SPARC:         "SUN SPARC",
	EM_386:           "Intel 80386",
	EM_68K:           "Motorola m68k family",
	EM_88K:           "Motorola m88k family",
	EM_IAMCU:         "Intel MCU",
	EM_860:           "Intel 80860",
	EM_MIPS:          "MIPS R3000 big-endian",
	EM_S370:          "IBM System/370",
	EM_MIPS_RS3_LE:   "MIPS R3000 little-endian",
	EM_PARISC:        "HPPA",
	EM_VPP500:        "Fujitsu VPP500",
	EM_SPARC32PLUS:   "Sun's \"v8plus\"",
	EM_960:           "Intel 80960",
	EM_PPC:           "PowerPC",
	EM_PPC64:         "PowerPC 64-bit",
	EM_S390:          "IBM S390",
	EM_SPU:           "IBM SPU/SPC",
	EM_V800:          "NEC V800 series",
	EM_FR20:          "Fujitsu FR20",
	EM_RH32:          "TRW RH-32",
	EM_RCE:           "Motorola RCE",
	EM_ARM:           "ARM",
	EM_FAKE_ALPHA:    "Digital Alpha",
	EM_SH:            "Hitachi SH",
	EM_SPARCV9:       "SPARC v9 64-bit",
	EM_TRICORE:       "Siemens Tricore",
	EM_ARC:           "Argonaut RISC Core",
	EM_H8_300:        "Hitachi H8/300",
	EM_H8_300H:       "Hitachi H8/300H",
	EM_H8S:           "Hitachi H8S",
	EM_H8_500:        "Hitachi H8/500",
	EM_IA_64:         "Intel Merced",
	EM_MIPS_X:        "Stanford MIPS-X",
	EM_COLDFIRE:      "Motorola Coldfire",
	EM_68HC12:        "Motorola M68HC12",
	EM_MMA:           "Fujitsu MMA Multimedia Accelerator",
	EM_PCP:           "Siemens PCP",
	EM_NCPU:          "Sony nCPU embeeded RISC",
	EM_NDR1:          "Denso NDR1 microprocessor",
	EM_STARCORE:      "Motorola Start*Core processor",
	EM_ME16:          "Toyota ME16 processor",
	EM_ST100:         "STMicroelectronic ST100 processor",
	EM_TINYJ:         "Advanced Logic Corp. Tinyj emb.fam",
	EM_X86_64:        "AMD x86-64 architecture",
	EM_PDSP:          "Sony DSP Processor",
	EM_PDP10:         "Digital PDP-10",
	EM_PDP11:         "Digital PDP-11",
	EM_FX66:          "Siemens FX66 microcontroller",
	EM_ST9PLUS:       "STMicroelectronics ST9+ 8/16 mc",
	EM_ST7:           "STmicroelectronics ST7 8 bit mc",
	EM_68HC16:        "Motorola MC68HC16 microcontroller",
	EM_68HC11:        "Motorola MC68HC11 microcontroller",
	EM_68HC08:        "Motorola MC68HC08 microcontroller",
	EM_68HC05:        "Motorola MC68HC05 microcontroller",
	EM_SVX:           "Silicon Graphics SVx",
	EM_ST19:          "STMicroelectronics ST19 8 bit mc",
	EM_VAX:           "Digital VAX",
	EM_CRIS:          "Axis Communications 32-bit emb.proc",
	EM_JAVELIN:       "Infineon Technologies 32-bit emb.proc",
	EM_FIREPATH:      "Element 14 64-bit DSP Processor",
	EM_ZSP:           "LSI Logic 16-bit DSP Processor",
	EM_MMIX:          "Donald Knuth's educational 64-bit proc",
	EM_HUANY:         "Harvard University machine-independent object files",
	EM_PRISM:         "SiTera Prism",
	EM_AVR:           "Atmel AVR 8-bit microcontroller",
	EM_FR30:          "Fujitsu FR30",
	EM_D10V:          "Mitsubishi D10V",
	EM_D30V:          "Mitsubishi D30V",
	EM_V850:          "NEC v850",
	EM_M32R:          "Mitsubishi M32R",
	EM_MN10300:       "Matsushita MN10300",
	EM_MN10200:       "Matsushita MN10200",
	EM_PJ:            "picoJava",
	EM_OPENRISC:      "OpenRISC 32-bit embedded processor",
	EM_ARC_COMPACT:   "ARC International ARCompact",
	EM_XTENSA:        "Tensilica Xtensa Architecture",
	EM_VIDEOCORE:     "Alphamosaic VideoCore",
	EM_TMM_GPP:       "Thompson Multimedia General Purpose Proc",
	EM_NS32K:         "National Semi. 32000",
	EM_TPC:           "Tenor Network TPC",
	EM_SNP1K:         "Trebia SNP 1000",
	EM_ST200:         "STMicroelectronics ST200",
	EM_IP2K:          "Ubicom IP2xxx",
	EM_MAX:           "MAX processor",
	EM_CR:            "National Semi. CompactRISC",
	EM_F2MC16:        "Fujitsu F2MC16",
	EM_MSP430:        "Texas Instruments msp430",
	EM_BLACKFIN:      "Analog Devices Blackfin DSP",
	EM_SE_C33:        "Seiko Epson S1C33 family",
	EM_SEP:           "Sharp embedded microprocessor",
	EM_ARCA:          "Arca RISC",
	EM_UNICORE:       "PKU-Unity & MPRC Peking Uni. mc series",
	EM_EXCESS:        "eXcess configurable cpu",
	EM_DXP:           "Icera Semi. Deep Execution Processor",
	EM_ALTERA_NIOS2:  "Altera Nios II",
	EM_CRX:           "National Semi. CompactRISC CRX",
	EM_XGATE:         "Motorola XGATE",
	EM_C166:          "Infineon C16x/XC16x",
	EM_M16C:          "Renesas M16C",
	EM_DSPIC30F:      "Microchip Technology dsPIC30F",
	EM_CE:            "Freescale Communication Engine RISC",
	EM_M32C:          "Renesas M32C",
	EM_TSK3000:       "Altium TSK3000",
	EM_RS08:          "Freescale RS08",
	EM_SHARC:         "Analog Devices SHARC family",
	EM_ECOG2:         "Cyan Technology eCOG2",
	EM_SCORE7:        "Sunplus S+core7 RISC",
	EM_DSP24:         "New Japan Radio (NJR) 24-bit DSP",
	EM_VIDEOCORE3:    "Broadcom VideoCore III",
	EM_LATTICEMICO32: "RISC for Lattice FPGA",
	EM_SE_C17:        "Seiko Epson C17",
	EM_TI_C6000:      "Texas Instruments TMS320C6000 DSP",
	EM_TI_C2000:      "Texas Instruments TMS320C2000 DSP",
	EM_TI_C5500:      "Texas Instruments TMS320C55x DSP",
	EM_TI_ARP32:      "Texas Instruments App. Specific RISC",
	EM_TI_PRU:        "Texas Instruments Prog. Realtime Unit",
	EM_MMDSP_PLUS:    "STMicroelectronics 64bit VLIW DSP",
	EM_CYPRESS_M8C:   "Cypress M8C",
	EM_R32C:          "Renesas R32C",
	EM_TRIMEDIA:      "NXP Semi. TriMedia",
	EM_QDSP6:         "QUALCOMM DSP6",
	EM_8051:          "Intel 8051 and variants",
	EM_STXP7X:        "STMicroelectronics STxP7x",
	EM_NDS32:         "Andes Tech. compact code emb. RISC",
	EM_ECOG1X:        "Cyan Technology eCOG1X",
	EM_MAXQ30:        "Dallas Semi. MAXQ30 mc",
	EM_XIMO16:        "New Japan Radio (NJR) 16-bit DSP",
	EM_MANIK:         "M2000 Reconfigurable RISC",
	EM_CRAYNV2:       "Cray NV2 vector architecture",
	EM_RX:            "Renesas RX",
	EM_METAG:         "Imagination Tech. META",
	EM_MCST_ELBRUS:   "MCST Elbrus",
	EM_ECOG16:        "Cyan Technology eCOG16",
	EM_CR16:          "National Semi. CompactRISC CR16",
	EM_ETPU:          "Freescale Extended Time Processing Unit",
	EM_SLE9X:         "Infineon Tech. SLE9X",
	EM_L10M:          "Intel L10M",
	EM_K10M:          "Intel K10M",
	EM_AARCH64:       "ARM AARCH64",
	EM_AVR32:         "Amtel 32-bit microprocessor",
	EM_STM8:          "STMicroelectronics STM8",
	EM_TILE64:        "Tilera TILE64",
	EM_TILEPRO:       "Tilera TILEPro",
	EM_MICROBLAZE:    "Xilinx MicroBlaze",
	EM_CUDA:          "NVIDIA CUDA",
	EM_TILEGX:        "Tilera TILE-Gx",
	EM_CLOUDSHIELD:   "CloudShield",
	EM_COREA_1ST:     "KIPO-KAIST Core-A 1st gen.",
	EM_COREA_2ND:     "KIPO-KAIST Core-A 2nd gen.",
	EM_ARCV2:         "Synopsys ARCv2 ISA. ",
	EM_OPEN8:         "Open8 RISC",
	EM_RL78:          "Renesas RL78",
	EM_VIDEOCORE5:    "Broadcom VideoCore V",
	EM_78KOR:         "Renesas 78KOR",
	EM_56800EX:       "Freescale 56800EX DSC",
	EM_BA1:           "Beyond BA1",
	EM_BA2:           "Beyond BA2",
	EM_XCORE:         "XMOS xCORE",
	EM_MCHP_PIC:      "Microchip 8-bit PIC(r)",
	EM_INTELGT:       "Intel Graphics Technology",
	EM_KM32:          "KM211 KM32",
	EM_KMX32:         "KM211 KMX32",
	EM_EMX16:         "KM211 KMX16",
	EM_EMX8:          "KM211 KMX8",
	EM_KVARC:         "KM211 KVARC",
	EM_CDP:           "Paneve CDP",
	EM_COGE:          "Cognitive Smart Memory Processor",
	EM_COOL:          "Bluechip CoolEngine",
	EM_NORC:          "Nanoradio Optimized RISC",
	EM_CSR_KALIMBA:   "CSR Kalimba",
	EM_Z80:           "Zilog Z80",
	EM_VISIUM:        "Controls and Data Services VISIUMcore",
	EM_FT32:          "FTDI Chip FT32",
	EM_MOXIE:         "Moxie processor",
	EM_AMDGPU:        "AMD GPU",
	EM_RISCV:         "RISC-V",
	EM_BPF:           "Linux BPF -- in-kernel virtual machine",
	EM_CSKY:          "C-SKY",
}

/* Program segment header.  */

/* Legal values for p_type (segment type).  */
const (
	PT_NULL         = 0          /* Program header table entry unused */
	PT_LOAD         = 1          /* Loadable program segment */
	PT_DYNAMIC      = 2          /* Dynamic linking information */
	PT_INTERP       = 3          /* Program interpreter */
	PT_NOTE         = 4          /* Auxiliary information */
	PT_SHLIB        = 5          /* Reserved */
	PT_PHDR         = 6          /* Entry for header table itself */
	PT_TLS          = 7          /* Thread-local storage segment */
	PT_NUM          = 8          /* Number of defined types */
	PT_LOOS         = 0x60000000 /* Start of OS-specific */
	PT_GNU_EH_FRAME = 0x6474e550 /* GCC .eh_frame_hdr segment */
	PT_GNU_STACK    = 0x6474e551 /* Indicates stack executability */
	PT_GNU_RELRO    = 0x6474e552 /* Read-only after relocation */
	PT_GNU_PROPERTY = 0x6474e553 /* GNU property */
	PT_LOSUNW       = 0x6ffffffa /* Reserved for Sun-specific semantics */
	PT_SUNWBSS      = 0x6ffffffa /* Sun Specific segment */
	PT_SUNWSTACK    = 0x6ffffffb /* Stack segment */
	PT_HISUNW       = 0x6fffffff /* Reserved for Sun-specific semantics */
	PT_HIOS         = 0x6fffffff /* End of OS-specific */
	PT_LOPROC       = 0x70000000 /* Start of processor-specific */

	PT_RISCV_ATTRIBUTES = 0x70000003 /* Location of RISC-V ELF attribute section */

	PT_HIPROC = 0x7fffffff /* End of processor-specific */
)

var p_type = map[Elf64_Word]string{
	PT_NULL:    "NULL",
	PT_LOAD:    "LOAD",
	PT_DYNAMIC: "DYNAMIC",
	PT_INTERP:  "INTERP",
	PT_NOTE:    "NOTE",
	PT_SHLIB:   "SHLIB",
	PT_PHDR:    "PHDR",
	PT_TLS:     "TLS",
	PT_NUM:     "NUM",

	PT_RISCV_ATTRIBUTES: "RISCV_ATTRIBUT",
}

/* Legal values for p_flags (segment flags).  */
const (
	PF_X        = (1 << 0)   /* Segment is executable */
	PF_W        = (1 << 1)   /* Segment is writable */
	PF_R        = (1 << 2)   /* Segment is readable */
	PF_MASKOS   = 0x0ff00000 /* OS-specific */
	PF_MASKPROC = 0xf0000000 /* Processor-specific */
)

func getSegmentFlags(segFlags Elf64_Word) string {
	flags := ""

	if segFlags&PF_R != 0 {
		flags += "R"
	}

	if segFlags&PF_W != 0 {
		flags += "W"
	}

	if segFlags&PF_X != 0 {
		flags += " E"
	}
	return flags
}

/* Section header */

/* Legal values for sh_type (section type).  */
const (
	SHT_NULL          = 0  /* Section header table entry unused */
	SHT_PROGBITS      = 1  /* Program specific (private) data */
	SHT_SYMTAB        = 2  /* Link editing symbol table */
	SHT_STRTAB        = 3  /* A string table */
	SHT_RELA          = 4  /* Relocation entries with addends */
	SHT_HASH          = 5  /* A symbol hash table */
	SHT_DYNAMIC       = 6  /* Information for dynamic linking */
	SHT_NOTE          = 7  /* Information that marks file */
	SHT_NOBITS        = 8  /* Section occupies no space in file */
	SHT_REL           = 9  /* Relocation entries, no addends */
	SHT_SHLIB         = 10 /* Reserved, unspecified semantics */
	SHT_DYNSYM        = 11 /* Dynamic linking symbol table */
	SHT_INIT_ARRAY    = 14 /* Array of ptrs to init functions */
	SHT_FINI_ARRAY    = 15 /* Array of ptrs to finish functions */
	SHT_PREINIT_ARRAY = 16 /* Array of ptrs to pre-init funcs */
	SHT_GROUP         = 17 /* Section contains a section group */
	SHT_SYMTAB_SHNDX  = 18 /* Indices for SHN_XINDEX entries */
	SHT_RELR          = 19 /* RELR relative relocations */
)

var sh_type = map[Elf64_Word]string{
	SHT_NULL:          "NULL",
	SHT_PROGBITS:      "PROGBITS",
	SHT_SYMTAB:        "SYMTAB",
	SHT_STRTAB:        "STRTAB",
	SHT_RELA:          "RELA",
	SHT_HASH:          "HASH",
	SHT_DYNAMIC:       "DYNAMIC",
	SHT_NOTE:          "NOTE",
	SHT_NOBITS:        "NOBITS",
	SHT_REL:           "REL",
	SHT_SHLIB:         "SHLIB",
	SHT_DYNSYM:        "DYNSYM",
	SHT_INIT_ARRAY:    "INIT_ARRAY",
	SHT_FINI_ARRAY:    "FINI_ARRAY",
	SHT_PREINIT_ARRAY: "PREINIT_ARRAY",
	SHT_GROUP:         "GROUP",
	SHT_SYMTAB_SHNDX:  "SYMTAB_SHNDX",
	SHT_RELR:          "RELR",
}

/* Legal values for sh_flags (section flags).  */
const (
	SHF_WRITE            = (1 << 0)  /* Writable */
	SHF_ALLOC            = (1 << 1)  /* Occupies memory during execution */
	SHF_EXECINSTR        = (1 << 2)  /* Executable */
	SHF_MERGE            = (1 << 4)  /* Might be merged */
	SHF_STRINGS          = (1 << 5)  /* Contains nul-terminated strings */
	SHF_INFO_LINK        = (1 << 6)  /* `sh_info' contains SHT index */
	SHF_LINK_ORDER       = (1 << 7)  /* Preserve order after combining */
	SHF_OS_NONCONFORMING = (1 << 8)  /* Non-standard OS specific handling required */
	SHF_GROUP            = (1 << 9)  /* Section is member of a group.  */
	SHF_TLS              = (1 << 10) /* Section hold thread-local data.  */
	SHF_COMPRESSED       = (1 << 11) /* Section with compressed data. */
)

func getSectionFlags(sh_flags Elf64_XWord) string {
	flags := ""

	if sh_flags&SHF_WRITE != 0 {
		flags += "W"
	}
	if sh_flags&SHF_ALLOC != 0 {
		flags += "A"
	}
	if sh_flags&SHF_EXECINSTR != 0 {
		flags += "X"
	}
	if sh_flags&SHF_MERGE != 0 {
		flags += "M"
	}
	if sh_flags&SHF_STRINGS != 0 {
		flags += "S"
	}
	if sh_flags&SHF_INFO_LINK != 0 {
		flags += "I"
	}
	if sh_flags&SHF_LINK_ORDER != 0 {
		flags += "L"
	}
	if sh_flags&SHF_OS_NONCONFORMING != 0 {
		flags += "O"
	}
	if sh_flags&SHF_GROUP != 0 {
		flags += "G"
	}
	if sh_flags&SHF_TLS != 0 {
		flags += "T"
	}
	if sh_flags&SHF_COMPRESSED != 0 {
		flags += "C"
	}
	return flags
}

/* Symbol Table */

/* Legal values for ST_BIND subfield of st_info (symbol binding).  */
type ST_BIND int

const (
	STB_LOCAL      = 0  /* Local symbol */
	STB_GLOBAL     = 1  /* Global symbol */
	STB_WEAK       = 2  /* Weak symbol */
	STB_NUM        = 3  /* Number of defined types.  */
	STB_LOOS       = 10 /* Start of OS-specific */
	STB_GNU_UNIQUE = 10 /* Unique symbol.  */
	STB_HIOS       = 12 /* End of OS-specific */
	STB_LOPROC     = 13 /* Start of processor-specific */
	STB_HIPROC     = 15 /* End of processor-specific */
)

var sym_bind = map[Elf_UChar]string{
	STB_LOCAL:  "LOCAL",
	STB_GLOBAL: "GLOBAL",
	STB_WEAK:   "WEAK",
}

/* Legal values for ST_TYPE subfield of st_info (symbol type).  */
const (
	STT_NOTYPE    = 0  /* Symbol type is unspecified */
	STT_OBJECT    = 1  /* Symbol is a data object */
	STT_FUNC      = 2  /* Symbol is a code object */
	STT_SECTION   = 3  /* Symbol associated with a section */
	STT_FILE      = 4  /* Symbol's name is file name */
	STT_COMMON    = 5  /* Symbol is a common data object */
	STT_TLS       = 6  /* Symbol is thread-local data object*/
	STT_NUM       = 7  /* Number of defined types.  */
	STT_RELC      = 8  /* Complex relocation expression */
	STT_SRELC     = 9  /* Signed Complex relocation expression */
	STT_LOOS      = 10 /* Start of OS-specific */
	STT_GNU_IFUNC = 10 /* Symbol is indirect code object */
	STT_HIOS      = 12 /* End of OS-specific */
	STT_LOPROC    = 13 /* Start of processor-specific */
	STT_HIPROC    = 15 /* End of processor-specific */
)

var sym_type = map[Elf_UChar]string{
	STT_NOTYPE:  "NOTYPE",
	STT_OBJECT:  "OBJECT",
	STT_FUNC:    "FUNC",
	STT_SECTION: "SECTION",
	STT_FILE:    "FILE",
	STT_COMMON:  "COMMON",
	STT_TLS:     "TLS",
	STT_RELC:    "RELC",
	STT_SRELC:   "SRELC",
}

/* Symbol visibility specification encoded in the st_other field.  */
const (
	STV_DEFAULT   = 0 /* Default symbol visibility rules */
	STV_INTERNAL  = 1 /* Processor specific hidden class */
	STV_HIDDEN    = 2 /* Sym unavailable in other modules */
	STV_PROTECTED = 3 /* Not preemptible, not exported */
)

var sym_vis = map[Elf_UChar]string{
	STV_DEFAULT:   "DEFAULT",
	STV_INTERNAL:  "INTERNAL",
	STV_HIDDEN:    "HIDDEN",
	STV_PROTECTED: "PROTECTED",
}

/* Special section indices.  */
const (
	SHN_UNDEF  = 0      /* Undefined section */
	SHN_ABS    = 0xfff1 /* Associated symbol is absolute */
	SHN_COMMON = 0xfff2 /* Associated symbol is common */
	SHN_XINDEX = 0xffff /* Index is in extra table.  */
)

var sym_idx = map[Elf64_Half]string{
	SHN_UNDEF:  "UND",
	SHN_ABS:    "ABS",
	SHN_COMMON: "COM",
	// SHN_XINDEX: "COM",
}
