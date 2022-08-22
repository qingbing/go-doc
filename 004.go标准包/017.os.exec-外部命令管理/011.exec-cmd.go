package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	var cmd *exec.Cmd
	// cmd.CombinedOutput()
	fmt.Println("=== cmd.CombinedOutput() ===")
	cmd = exec.Command("ls", "-al")
	fmt.Println("最终执行的命令", cmd.String())
	bs, _ := cmd.CombinedOutput()
	fmt.Println(string(bs))

	// cmd.Output()
	fmt.Println("=== cmd.Output() ===")
	cmd = exec.Command("ls", "-al")
	fmt.Println("最终执行的命令", cmd.String())
	bs, _ = cmd.Output()
	fmt.Println(string(bs))

	// cmd.Start()
	//fmt.Println("=== cmd.Start() ===")
	//cmd = exec.Command("ping", "baidu.com")
	//cmd.Stdout = os.Stdout
	//cmd.Start()
	//cmd.Wait()

	// cmd.Run()
	//fmt.Println("=== cmd.Run() ===")
	//cmd = exec.Command("ping", "baidu.com")
	//fmt.Println("最终执行的命令", cmd.String())
	//cmd.Stdout = os.Stdout
	//cmd.Run()

	// cmd.StdoutPipe()
	fmt.Println("=== cmd.StdoutPipe() ===")
	cmd = exec.Command("ls", "-al")
	rStdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	bs, _ = io.ReadAll(rStdout)
	fmt.Printf("%s\n", bs)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
