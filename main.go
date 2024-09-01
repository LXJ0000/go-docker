package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Unknown command")
	}
}

func run() {
	cmd := exec.Command(os.Args[2])
	
	cmd.SysProcAttr = &syscall.SysProcAttr{
	    
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
