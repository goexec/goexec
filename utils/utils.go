package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/goexec/goexec/constants"
)

// RunGoProgram executes a go program.
func RunGoProgram(args ...string) error {
	// go run
	argv := []string{"run"}
	argv = append(argv, args...)

	stdout, stderr, err := execShellCmd("go", argv...)

	if stderr != "" {
		fmt.Println(stderr)
	}
	fmt.Println(stdout)
	if err != nil {
		return err
	}

	return nil
}

// RunCProgram executes a c program.
func RunCProgram(args ...string) error {
	tempPath := filepath.Join(constants.TempDir, "a.out")
	if runtime.GOOS == "windows" {
		tempPath = "out.exe"
	}

	// gcc
	argv := []string{"-o", tempPath}
	argv = append(argv, args[0])

	stdout, stderr, err := execShellCmd("gcc", argv...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
	if err != nil {
		return err
	}

	stdout, stderr, err = execShellCmd(tempPath, args[1:]...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	fmt.Println(stdout)
	if err != nil {
		return err
	}

	_, stderr, _ = execShellCmd("rm", tempPath)
	if stderr != "" {
		fmt.Println(stderr)
	}

	return nil
}

// RunCPPProgram executes a c program.
func RunCPPProgram(args ...string) error {
	tempPath := filepath.Join(constants.TempDir, "a.out")
	if runtime.GOOS == "windows" {
		tempPath = "out.exe"
	}

	// g++
	argv := []string{"-o", tempPath}
	argv = append(argv, args[0])

	stdout, stderr, err := execShellCmd("g++", argv...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
	if err != nil {
		return err
	}

	stdout, stderr, err = execShellCmd(tempPath, args[1:]...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	fmt.Println(stdout)
	if err != nil {
		return err
	}

	_, stderr, _ = execShellCmd("rm", tempPath)
	if stderr != "" {
		fmt.Println(stderr)
	}

	return nil
}

// RunJavaProgram executes a java program.
func RunJavaProgram(args ...string) error {
	fileName := FileNameWithoutExtension(args[0])
	className := fileName
	if strings.ContainsAny(className, "/") {
		className = strings.ReplaceAll(className, "/", ".")
	}
	stdout, stderr, err := execShellCmd("javac", args[0])
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
	if err != nil {
		return err
	}
	argv := []string{className}
	argv = append(argv, args[1:]...)
	stdout, stderr, err = execShellCmd("java", argv...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	fmt.Println(stdout)
	if err != nil {
		return err
	}
	_, stderr, _ = execShellCmd("rm", fileName+".class")
	if stderr != "" {
		fmt.Println(stderr)
	}

	return nil
}

// RunRubyProgram executes a ruby program.
func RunRubyProgram(args ...string) error {
	stdout, stderr, err := execShellCmd("ruby", args...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
	if err != nil {
		return err
	}
	return nil
}

// RunLuaProgram executes a lua program.
func RunLuaProgram(args ...string) error {
	stdout, stderr, err := execShellCmd("lua", args...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
	if err != nil {
		return err
	}
	return nil
}

// RunPythonProgram executes a python program.
func RunPythonProgram(args ...string) error {
	stdout, stderr, err := execShellCmd("python3", args...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
	if err != nil {
		return err
	}
	return nil
}

// RunPerlProgram executes a perl program.
func RunPerlProgram(args ...string) error {
	stdout, stderr, err := execShellCmd("perl", args...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
	if err != nil {
		return err
	}
	return nil
}

// RunJSProgram executes a javascript program.
func RunJSProgram(args ...string) error {
	stdout, stderr, err := execShellCmd("node", args...)
	if stderr != "" {
		fmt.Println(stderr)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
	if err != nil {
		return err
	}
	return nil
}

// execShellCmd executes a shell command and returns the output.
func execShellCmd(app string, args ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	if !commandExists(app) {
		return "", "", fmt.Errorf("command not found: %s", app)
	}

	cmd := exec.Command(app, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", "", err
	}
	return stdout.String(), stderr.String(), nil
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// FileNameWithoutExtension returns the file name without extension.
func FileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

// GetLangName returns lang name by extension
func GetLangName(filename string) string {
	filename = strings.ToLower(filename)
	if strings.HasSuffix(filename, ".go") {
		return "go"
	}
	if strings.HasSuffix(filename, ".c") {
		return "c"
	}
	if strings.HasSuffix(filename, ".cpp") || strings.HasSuffix(filename, ".cc") {
		return "cpp"
	}
	if strings.HasSuffix(filename, ".cs") {
		return "c-sharp"
	}
	if strings.HasSuffix(filename, ".java") {
		return "java"
	}
	if strings.HasSuffix(filename, ".rb") {
		return "ruby"
	}
	if strings.HasSuffix(filename, ".py") {
		return "python"
	}
	if strings.HasSuffix(filename, ".lua") {
		return "lua"
	}
	if strings.HasSuffix(filename, ".pl") {
		return "perl"
	}
	if strings.HasSuffix(filename, ".php") {
		return "php"
	}
	if strings.HasSuffix(filename, ".rs") {
		return "rust"
	}
	if strings.HasSuffix(filename, ".dart") {
		return "dart"
	}
	if strings.HasSuffix(filename, ".js") {
		return "javascript"
	}
	if strings.HasSuffix(filename, ".ts") {
		return "typescript"
	}

	return "unknown"
}