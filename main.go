package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/goexec/goexec/utils"
)

var (
	flagHelp    bool
	flagVersion bool
)

func init() {
	flag.BoolVar(&flagHelp, "h", false, "Show this help")
	flag.BoolVar(&flagHelp, "help", false, "Show this help")
	flag.BoolVar(&flagVersion, "v", false, "Show version")
	flag.BoolVar(&flagVersion, "version", false, "Show version")
}

func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NArg() == 0 {
		printUsage()
	}

	// if user wants to run a command, run it
	if flag.NArg() > 0 {
		runCommand(flag.Args())
	}

	// if user wants help, print usage
	if flagHelp {
		printUsage()
	}

	// if user wants version, print version
	if flagVersion {
		printVersion()
	}

}

func runCommand(args []string) {
	fmt.Println("Running command: " + strings.Join(args, " "))
	lang := utils.GetLangName(args[0])

	switch lang {
	case "go":
		err := utils.RunGoProgram(args...)
		if err != nil {
			fmt.Println(err)
		}

	case "c":
		err := utils.RunCProgram(args...)
		if err != nil {
			fmt.Println(err)
		}
	case "cpp":
		err := utils.RunCPPProgram(args...)
		if err != nil {
			fmt.Println(err)
		}

	case "java":
		err := utils.RunJavaProgram(args...)
		if err != nil {
			fmt.Println(err)
		}
	case "ruby":
		err := utils.RunRubyProgram(args...)
		if err != nil {
			fmt.Println(err)
		}
	case "lua":
		err := utils.RunLuaProgram(args...)
		if err != nil {
			fmt.Println(err)
		}
	case "python":
		err := utils.RunPythonProgram(args...)
		if err != nil {
			fmt.Println(err)
		}
	case "perl":
		err := utils.RunPerlProgram(args...)
		if err != nil {
			fmt.Println(err)
		}
	case "javascript":
		err := utils.RunJSProgram(args...)
		if err != nil {
			fmt.Println(err)
		}

	case "unknown":
		fmt.Println("Unknown language")

	default:
		fmt.Println("this language is not yet supported or invalid file format")
	}
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}

func printVersion() {
	version := "1.0.0"
	fmt.Println(os.Args[0] + " version " + version)
}
