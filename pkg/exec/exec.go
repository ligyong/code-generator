package exec

import (
	"fmt"
	"os"
	"os/exec"
)

type Exec interface {
	Run() error
}

type command struct {
	cmd *exec.Cmd
}

func New(cmdStr string) Exec {
	//fmt.Println("******************************************************")
	//fmt.Println(cmdStr)
	//fmt.Println("******************************************************")

	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin

	return command{cmd: cmd}
}

func (c command) Run() error {
	if err := c.cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return nil
}
