package main

import (
	"fmt"
	"os"
	"strings"
)

var options = map[string]bool{
	"header":   false,
	"sections": false,
	"segments": false,
	"symbols":  false,
	"all":      false,
	"help":     true,
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	paths, err := handleArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		printUsage()
		return
	}

	if err = parseFiles(paths); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		printUsage()
		return
	}
}

func parseFiles(paths []string) error {
	if len(paths) == 0 {
		return fmt.Errorf("elfparser: Warning: Nothing to do")
	}

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		parser, err := LoadData(file)
		if err != nil {
			return err
		}
		if options["all"] {
			parser.PrintEhdr()
			parser.PrintShdrs()
			parser.PrintPhdrs()
			parser.PrintSyms()
			options["help"] = false
			continue
		}
		if options["header"] {
			parser.PrintEhdr()
			options["help"] = false
		}
		if options["sections"] {
			parser.PrintShdrs()
			options["help"] = false
		}
		if options["segments"] {
			parser.PrintPhdrs()
			options["help"] = false
		}

		if options["symbols"] {
			parser.PrintSyms()
			options["help"] = false
		}
		if options["help"] {
			printUsage()
		}
	}
	return nil
}

func handleArgs(args []string) ([]string, error) {
	paths := []string{}
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			switch arg {
			case "--all":
				options["all"] = true
			case "--file-header":
				options["header"] = true
			case "--segments":
				options["segments"] = true
			case "--sections":
				options["sections"] = true
			case "--symbols":
				options["symbols"] = true
			case "--help":
				options["help"] = true
			default:
				return paths, fmt.Errorf("elfparser: unrecognized option: %s", arg)
			}
		} else if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-a":
				options["all"] = true
			case "-h":
				options["header"] = true
			case "-l":
				options["segments"] = true
			case "-S":
				options["sections"] = true
			case "-s":
				options["symbols"] = true
			case "-H":
				options["help"] = true
			default:
				return paths, fmt.Errorf("elfparser: unrecognized option: %s", arg)
			}
		} else {
			paths = append(paths, arg)
		}
	}

	return paths, nil
}

func printUsage() {
	var usage = `Usage: parser <option(s)> [executable]
  Display information about the contents of ELF format files
  Options are:
  -a --all          equivalent to: -h -l -S -s
  -h --file-header  Display the Elf file header
  -l --segments     Display the program headers
  -S --sections     Display the sections' header
  -s --symbols      Display the symbol table
  -H --help         Display this information`
	fmt.Println(usage)
}
