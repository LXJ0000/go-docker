package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Printf("Process => %v [%d]\n", os.Args, os.Getpid())
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Unknown command")
	}
}

func run() {
	cmd := exec.Command(os.Args[0], append([]string{"child"}, os.Args[2])...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {
	cmd := exec.Command(os.Args[2]) // 避免对主进程的 hostname 进行修改
	syscall.Sethostname([]byte("container"))
	// MS_NOEXEC 本文件系统不允许运行其他程序
	// MS_NOSUID 本文件系统运行其他程序的时候，不允许set-user-id和set-group-id
	// MS_NODEV 从Linux2.4以来，所有mount的系统都会默认设定的参数
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "") // TODO 隔离文件系统
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	syscall.Unmount("/proc", 0)
}
