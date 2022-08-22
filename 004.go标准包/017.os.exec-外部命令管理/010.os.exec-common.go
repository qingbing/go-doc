package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	var cmd *exec.Cmd

	// LookPath()
	fmt.Println("======== LookPath ========")
	if path, err := exec.LookPath("pwd"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("(pwd) path (%v)\n", path)
	}

	// exec.Command()
	fmt.Println("======== exec.Command ========")
	cmd = exec.Command("ls", "-al")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	} else {
		out.WriteTo(os.Stdout)
		//fmt.Println(out.String())
	}

	// exec.Cmd
	fmt.Println("======== exec.Cmd ========")
	cmd = &exec.Cmd{
		Path:   "/bin/ls",
		Args:   []string{"ls", "-al"},
		Stdout: os.Stdout,
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	// exec.CommandContext
	fmt.Println("======== exec.CommandContext ========")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	cmd = exec.CommandContext(ctx, "ping", "baidu.com")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
