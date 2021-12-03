package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Exec(cmd string) bool {
	list := strings.Split(cmd, " ")
	cmdline := exec.Command(list[0], list[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmdline.Stdout = &out
	cmdline.Stderr = &stderr
	err := cmdline.Run()
	if err != nil {
		fmt.Println(stderr.String())
		return false
	} else {
		fmt.Println(out.String())
		return true
	}
}
