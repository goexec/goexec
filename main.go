package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goexec/goexec/utils"
)

var (
	// AppName is the name of the application
	AppName = "goexec"
	// Version is the version of the application
	Version = "1.0.0"
	// Commit is the commit that was used to build the application
	GitCommit = "none"
	// BuildDate is the date the application was built
	BuildDate = "unknown"
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

	// if user wants help, print usage
	if flagHelp {
		printUsage()
	}

	// if user wants version, print version
	if flagVersion {
		printVersion()
	}

	// if user does not supply flags, print usage
	if flag.NArg() == 0 {
		printUsage()
	}

	// if user wants to run a command, run it
	runCommand(flag.Args())

}

func runCommand(args []string) {
	// fmt.Println("Running command: " + strings.Join(args, " "))
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
	case "typescript":
		err := utils.RunTSProgram(args...)
		if err != nil {
			fmt.Println(err)
		}
	case "shell":
		err := utils.RunShellProgram(args...)
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
	shortCommit := shortGitCommit(GitCommit)
	version := fmt.Sprintf(os.Args[0]+" version: %s %s %s", Version, shortCommit, BuildDate)
	fmt.Println(version)
	os.Exit(0)
}

// shortGitCommit returns the short form of the git commit hash
func shortGitCommit(fullGitCommit string) string {
	shortCommit := ""
	if len(fullGitCommit) >= 7 {
		shortCommit = fullGitCommit[0:7]
	}

	return shortCommit
}
