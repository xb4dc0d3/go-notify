package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Tmux struct {
}

func (t Tmux) Attach(session_name, executable, path string) {
	cmd := exec.Command("tmux", "new", "-s", session_name, "-d", executable, "attach", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()

	if e != nil {
		fmt.Println(e)
	}

}
