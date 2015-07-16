package goexec

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os/exec"
)

func ShellExec(cmd *exec.Cmd) []byte {
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	return stdout
}

func Pipeline(cmd *exec.Cmd, writer []byte) []byte {
	r, w := io.Pipe()
	cmd.Stdin = r

	var buf bytes.Buffer
	cmd.Stdout = &buf

	cmd.Start()
	w.Write(writer)
	w.Close()
	cmd.Wait()
	return buf.Bytes()
}
