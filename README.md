# elf

A parser for ELF64 object files in Go.

The implementation primarily referenced TIS1.1.pdf[^1], with the ELF Format Cheatsheet[^2] providing conceptual clarity. 

I also benefited from finixbit/elf-parser[^3] and macmade/ELFDump[^4]'s implementations, and borrowed numerous constant definitions from elf/common.h[^5].

The tool mimics readelf[^6]'s command-line argument design and output formatting, maintaining behavioral consistency while implementing a minimal subset of its functionality.

## Usage

```
Usage: parser <option(s)> [executable]
  Display information about the contents of ELF format files
  Options are:
  -a --all          equivalent to: -h -l -S -s
  -h --file-header  Display the Elf file header
  -l --segments     Display the program headers
  -S --sections     Display the sections' header
  -s --symbols      Display the symbol table
  -H --help         Display this information
```

`./parser -h /usr/bin/ls`
<details>
  <summary>Output:</summary>
  
```
ELF Header:
  Magic:                                  7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                                  64 bits
  Data:                                   LSB
  Version:                                1
  OS/ABI:                                 UNIX System V ABI
  ABI version:                            0
  Byte index:                             0
  Type:                                   DYN
  Machine:                                AMD x86-64 architecture
  Version:                                0x1
  Entry point address:                    0x6aa0
  Program header offset:                  64
  Section header offset:                  136232
  Flags:                                  0
  Size of this header                     64 (bytes)
  Size of program headers                 56 (bytes)
  Number of program headers               13
  Size of section headers                 64 (bytes)
  Number of section headers:              31
  Section header string table index:      30
```
</details>

`./parser -l /usr/bin/ls`
<details>
  <summary>Output:</summary>
  
```

Elf file type is DYN
Entry point 0x6aa0
There are 13 program headers, starting at offset 64

Program Headers:
  Type           Offset             VirtAddr           PhysAddr
                 FileSiz            MemSiz              Flags  Align
  PHDR           0x0000000000000040 0x0000000000000040 0x0000000000000040
                 0x00000000000002d8 0x00000000000002d8  R        0x8
  INTERP         0x0000000000000318 0x0000000000000318 0x0000000000000318
                 0x000000000000001c 0x000000000000001c  R        0x1
  LOAD           0x0000000000000000 0x0000000000000000 0x0000000000000000
                 0x0000000000003458 0x0000000000003458  R        0x1000
  LOAD           0x0000000000004000 0x0000000000004000 0x0000000000004000
                 0x0000000000013091 0x0000000000013091  R E      0x1000
  LOAD           0x0000000000018000 0x0000000000018000 0x0000000000018000
                 0x0000000000007458 0x0000000000007458  R        0x1000
  LOAD           0x000000000001ffd0 0x0000000000020fd0 0x0000000000020fd0
                 0x00000000000012a8 0x0000000000002570  RW       0x1000
  DYNAMIC        0x0000000000020a58 0x0000000000021a58 0x0000000000021a58
                 0x0000000000000200 0x0000000000000200  RW       0x8
  NOTE           0x0000000000000338 0x0000000000000338 0x0000000000000338
                 0x0000000000000030 0x0000000000000030  R        0x8
  NOTE           0x0000000000000368 0x0000000000000368 0x0000000000000368
                 0x0000000000000044 0x0000000000000044  R        0x4
                 0x0000000000000338 0x0000000000000338 0x0000000000000338
                 0x0000000000000030 0x0000000000000030  R        0x8
                 0x000000000001cdcc 0x000000000001cdcc 0x000000000001cdcc
                 0x000000000000056c 0x000000000000056c  R        0x4
                 0x0000000000000000 0x0000000000000000 0x0000000000000000
                 0x0000000000000000 0x0000000000000000  RW       0x10
                 0x000000000001ffd0 0x0000000000020fd0 0x0000000000020fd0
                 0x0000000000001030 0x0000000000001030  R        0x1
```
</details>

`./parser -S /usr/bin/ls`
<details>
  <summary>Output:</summary>
  
```

There are 31 section headers, starting at offset 0x21428:

Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [ 0]                   NULL             0000000000000000  00000000
       0000000000000000  0000000000000000           0     0     0
  [ 1] .interp           PROGBITS         0000000000000318  00000318
       000000000000001c  0000000000000000   A       0     0     1
  [ 2] .note.gnu.propertyNOTE             0000000000000338  00000338
       0000000000000030  0000000000000000   A       0     0     8
  [ 3] .note.gnu.build-idNOTE             0000000000000368  00000368
       0000000000000024  0000000000000000   A       0     0     4
  [ 4] .note.ABI-tag     NOTE             000000000000038c  0000038c
       0000000000000020  0000000000000000   A       0     0     4
  [ 5] .gnu.hash                          00000000000003b0  000003b0
       000000000000004c  0000000000000000   A       6     0     8
  [ 6] .dynsym           DYNSYM           0000000000000400  00000400
       0000000000000b88  0000000000000018   A       7     1     8
  [ 7] .dynstr           STRTAB           0000000000000f88  00000f88
       00000000000005a6  0000000000000000   A       0     0     1
  [ 8] .gnu.version                       000000000000152e  0000152e
       00000000000000f6  0000000000000002   A       6     0     2
  [ 9] .gnu.version_r                     0000000000001628  00001628
       00000000000000c0  0000000000000000   A       7     2     8
  [10] .rela.dyn         RELA             00000000000016e8  000016e8
       0000000000001410  0000000000000018   A       6     0     8
  [11] .rela.plt         RELA             0000000000002af8  00002af8
       0000000000000960  0000000000000018  AI       6    25     8
  [12] .init             PROGBITS         0000000000004000  00004000
       000000000000001b  0000000000000000  AX       0     0     4
  [13] .plt              PROGBITS         0000000000004020  00004020
       0000000000000650  0000000000000010  AX       0     0    16
  [14] .plt.got          PROGBITS         0000000000004670  00004670
       0000000000000030  0000000000000010  AX       0     0    16
  [15] .plt.sec          PROGBITS         00000000000046a0  000046a0
       0000000000000640  0000000000000010  AX       0     0    16
  [16] .text             PROGBITS         0000000000004ce0  00004ce0
       00000000000123a2  0000000000000000  AX       0     0    16
  [17] .fini             PROGBITS         0000000000017084  00017084
       000000000000000d  0000000000000000  AX       0     0     4
  [18] .rodata           PROGBITS         0000000000018000  00018000
       0000000000004dcc  0000000000000000   A       0     0    32
  [19] .eh_frame_hdr     PROGBITS         000000000001cdcc  0001cdcc
       000000000000056c  0000000000000000   A       0     0     4
  [20] .eh_frame         PROGBITS         000000000001d338  0001d338
       0000000000002120  0000000000000000   A       0     0     8
  [21] .init_array       INIT_ARRAY       0000000000020fd0  0001ffd0
       0000000000000008  0000000000000008  WA       0     0     8
  [22] .fini_array       FINI_ARRAY       0000000000020fd8  0001ffd8
       0000000000000008  0000000000000008  WA       0     0     8
  [23] .data.rel.ro      PROGBITS         0000000000020fe0  0001ffe0
       0000000000000a78  0000000000000000  WA       0     0    32
  [24] .dynamic          DYNAMIC          0000000000021a58  00020a58
       0000000000000200  0000000000000010  WA       7     0     8
  [25] .got              PROGBITS         0000000000021c58  00020c58
       00000000000003a0  0000000000000008  WA       0     0     8
  [26] .data             PROGBITS         0000000000022000  00021000
       0000000000000278  0000000000000000  WA       0     0    32
  [27] .bss              NOBITS           0000000000022280  00021278
       00000000000012c0  0000000000000000  WA       0     0    32
  [28] .gnu_debugaltlink PROGBITS         0000000000000000  00021278
       0000000000000049  0000000000000000           0     0     1
  [29] .gnu_debuglink    PROGBITS         0000000000000000  000212c4
       0000000000000034  0000000000000000           0     0     4
  [30] .shstrtab         STRTAB           0000000000000000  000212f8
       000000000000012f  0000000000000000           0     0     1
```
</details>

`./parser -s /usr/bin/ls`
<details>
  <summary>Output:</summary>
  
```

Symbol table '.symtab' contains 123 entries:
   Num:    Value          Size Type    Bind   Vis      Ndx Name
     0: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT  UND
     1: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __ctype_toupper_loc
     2: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND getenv
     3: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND __progname
     4: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND sigprocmask
     5: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __snprintf_chk
     6: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND raise
     7: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __libc_start_main
     8: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND abort
     9: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __errno_location
    10: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strncmp
    11: 0000000000000000     0 NOTYPE  WEAK   DEFAULT  UND _ITM_deregisterTMCloneTable
    12: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND stdout
    13: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND localtime_r
    14: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND _exit
    15: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strcpy
    16: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __mbstowcs_chk
    17: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __fpending
    18: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND isatty
    19: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND sigaction
    20: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND iswcntrl
    21: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND wcswidth
    22: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND localeconv
    23: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND mbstowcs
    24: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND readlink
    25: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND clock_gettime
    26: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND setenv
    27: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND textdomain
    28: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fclose
    29: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND optind
    30: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND opendir
    31: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND getpwuid
    32: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND bindtextdomain
    33: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND dcgettext
    34: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __ctype_get_mb_cur_max
    35: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strlen
    36: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __stack_chk_fail
    37: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND getopt_long
    38: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND mbrtowc
    39: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND freecon
    40: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strchr
    41: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND getgrgid
    42: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND snprintf
    43: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __overflow
    44: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strrchr
    45: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND gmtime_r
    46: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND lseek
    47: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __assert_fail
    48: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fnmatch
    49: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND memset
    50: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND ioctl
    51: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND getcwd
    52: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND closedir
    53: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND lstat
    54: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND memcmp
    55: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND _setjmp
    56: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fputs_unlocked
    57: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND calloc
    58: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strcmp
    59: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND signal
    60: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND dirfd
    61: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fputc_unlocked
    62: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND optarg
    63: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __memcpy_chk
    64: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND sigemptyset
    65: 0000000000000000     0 NOTYPE  WEAK   DEFAULT  UND __gmon_start__
    66: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND memcpy
    67: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND program_invocation_name
    68: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND tzset
    69: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fileno
    70: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND tcgetpgrp
    71: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND readdir
    72: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND wcwidth
    73: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fflush
    74: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND nl_langinfo
    75: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strcoll
    76: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND mktime
    77: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __freading
    78: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fwrite_unlocked
    79: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND realloc
    80: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND stpncpy
    81: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND setlocale
    82: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __printf_chk
    83: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND statx
    84: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND timegm
    85: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strftime
    86: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND mempcpy
    87: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND memmove
    88: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND error
    89: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND __progname_full
    90: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fseeko
    91: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND strtoumax
    92: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND unsetenv
    93: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __cxa_atexit
    94: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND wcstombs
    95: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND getxattr
    96: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND gethostname
    97: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND sigismember
    98: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND exit
    99: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fwrite
   100: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __fprintf_chk
   101: 0000000000000000     0 NOTYPE  WEAK   DEFAULT  UND _ITM_registerTMCloneTable
   102: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND getfilecon
   103: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND fflush_unlocked
   104: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND mbsinit
   105: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND lgetfilecon
   106: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND program_invocation_short_name
   107: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND iswprint
   108: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND sigaddset
   109: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __ctype_tolower_loc
   110: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __ctype_b_loc
   111: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND stderr
   112: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND __sprintf_chk
   113: 00000000000220a0     8 OBJECT  GLOBAL DEFAULT   26 obstack_alloc_failed_handler
   114: 000000000000fc60   296 FUNC    GLOBAL DEFAULT   16 _obstack_newchunk
   115: 000000000000fc40    25 FUNC    GLOBAL DEFAULT   16 _obstack_begin_1
   116: 0000000000010680    55 FUNC    GLOBAL DEFAULT   16 _obstack_allocated_p
   117: 0000000000000000     0 FUNC    WEAK   DEFAULT  UND __cxa_finalize
   118: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND free
   119: 000000000000fc20    21 FUNC    GLOBAL DEFAULT   16 _obstack_begin
   120: 0000000000000000     0 FUNC    GLOBAL DEFAULT  UND malloc
   121: 0000000000010750    38 FUNC    GLOBAL DEFAULT   16 _obstack_memory_used
   122: 00000000000106c0   133 FUNC    GLOBAL DEFAULT   16 _obstack_free
```
</details>

# Reference

[^1]: [TIS1.1.pdf](https://refspecs.linuxfoundation.org/elf/TIS1.1.pdf)

[^2]: [ELF Format Cheatsheet](https://gist.github.com/x0nu11byt3/bcb35c3de461e5fb66173071a2379779)

[^3]: [finixbit/elf-parser](https://github.com/finixbit/elf-parser/)

[^4]: [macmade/ELFDump](https://github.com/macmade/ELFDump/)

[^5]: [elf/common.h](https://github.com/bminor/binutils-gdb/blob/master/include/elf/common.h)

[^6]: [readelf.c](https://github.com/bminor/binutils-gdb/blob/master/binutils/readelf.c)

