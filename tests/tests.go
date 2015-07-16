package main

import (
	"fmt"
	"lib/goexec"
	"os/exec"
)

func main() {

	func() {
		str := `top -b infinity | sort -s -n -k 1,1`
		cmd := exec.Command("sh", "-c", str)
		result := goexec.ShellExec(cmd)
		fmt.Printf("%s", result)
	}()

	func() {
		str := `cat /dev/stdin`
		cmd := exec.Command("sh", "-c", str)
		result := goexec.Pipeline(cmd, []byte("I CAN WRITE STUFF TO EXECED FUNCTIONS!\n"))
		fmt.Printf("%s", result)
	}()
}
