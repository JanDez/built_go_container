//build linux

package main;

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//Docker 			run image 	<cmd> <params>
// go run main/go	run			<cmd> <params>


func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("invalid command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CloneFlags: syscall.CLONE_NEWUTS
	}
	cmd.Run()
}

func child() {
	fmt.Printf("Running %v\n", os.Args[2:])

	syscall.Sethostname([]byte("docker"))

	cmd := exec.Command((os.Args[2]), os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}