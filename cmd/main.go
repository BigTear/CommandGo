package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/BigTear/CommandGo/internal/app"
)

func main() {
	app.Init()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {

	// 预处理命令
	switch runtime.GOOS {
	case "linux":
		input = strings.TrimSuffix(input, "\n")
	case "windows":
		input = strings.TrimSuffix(input, "\r\n")
	}

	// 分析命令参数
	args := strings.Split(input, " ")

	// 执行内置命令
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("请输入一个有效的路径。")
		}
		return os.Chdir(args[1])
	case "exit":
		interExit(0)
	case "quit":
		interExit(0)
	}

	// 调用exec执行命令
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()

}

func interExit(code int) {
	os.Exit(code)
}
