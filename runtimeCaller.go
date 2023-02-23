package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

/*
runtime.Caller() is a function in Go's standard library that is used to
retrieve information about the calling function's stack frame.
It returns the file name, line number, and function name of the specified call frame in the call stack.

The function takes an integer argument skip that specifies the number of stack frames to
skip before returning information about the calling function's stack frame.
A value of 0 for skip refers to the immediate caller of runtime.Caller().
*/
func main() {
	callerInfo := make([]uintptr, 1)
	n := runtime.Callers(1, callerInfo)
	if n == 0 {
		fmt.Println("Could not retrieve caller information")
	}
	caller := runtime.FuncForPC(callerInfo[0] - 1)
	if caller == nil {
		fmt.Println("Could not retrieve caller function name")
	}
	file, line := caller.FileLine(callerInfo[0] - 1)
	fmt.Printf("Caller function name: %s\n", caller.Name())
	fmt.Printf("Caller file name: %s\n", file)
	fmt.Printf("Caller line number: %d\n", line)

	fileDirPath := path.Dir(file)
	newFilePath := path.Join(fileDirPath, "race.go")
	if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
		fmt.Println("file not exist")
	}
	fmt.Println(newFilePath)
	run(newFilePath)
}

func run(file string) {
	if strings.HasSuffix(file, ".go") != true {
		fmt.Println("The file is not a go file")
		return
	}
	// Command to run another Go file
	cmd := exec.Command("go", "run", file)

	// Set the current working directory for the command
	//cmd.Dir = "/Users/nilei/GolandProjects/go-base"

	// Set the environment variables for the command
	//cmd.Env = append(os.Environ(), "VAR1=value1", "VAR2=value2")

	// Redirect the command's standard output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and wait for it to finish
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		return
	}
}
