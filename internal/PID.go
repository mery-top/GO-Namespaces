package internal


import(
	"fmt"
	"os/exec"
	"syscall"
	"os"
)


func PID(){
	cmd:= exec.Command("/bin/bash")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Launching shell in new PID namespace...")
	if err:= cmd.Run(); err!= nil{
		fmt.Println("Error", err)
	}

}

