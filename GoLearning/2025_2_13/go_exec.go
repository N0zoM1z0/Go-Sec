package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("id")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", out)
}
