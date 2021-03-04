package main

import "os/exec"
import . "fmt"

func main() {
	app := "git"

	arg0 := "status"

	cmdGit := exec.Command(app, arg0)
	stdout, err := cmdGit.Output()

	if err != nil {
		Println(err.Error())
		return
	}

	Print(string(stdout))
}
