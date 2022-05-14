package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	PATH_EXE string = "/bin/CommandGo.exe"
	PATH_SRC string = "/cmd/CommandGo/main.go"
	PATH_UPX string = "/tools/upx.exe"
)

var pathBase = ""

var pathSrc, pathExe, pathUpx string

func main() {
	args := (os.Args)
	cmd := "default"
	if len(args) > 1 {
		// 小写命令
		cmd = strings.ToLower(args[1])
		// 去除命令前的 - 或 --
		if cmd[0] == '-' {
			cmd = cmd[1:]
		}
		if cmd[0] == '-' {
			cmd = cmd[1:]
		}
	}
	// 获取根路径
	if pathBase == "" {
		cwd, err := os.Getwd()
		if err != nil {
			println("取当前目录失败！", err)
		}
		pathBase = cwd
	}
	// 获取目录
	pathSrc = filepath.Join(pathBase, PATH_SRC)
	pathExe = filepath.Join(pathBase, PATH_EXE)
	pathUpx = filepath.Join(pathBase, PATH_UPX)

	// 执行命令
	switch cmd {
	case "build":
		Build()
	case "b":
		Build()
	case "run":
		Run()
	case "r":
		Run()
	case "upx":
		UPX()
	case "u":
		UPX()
	case "clean":
		Clean()
	case "c":
		Clean()
	case "help":
		Help()
	case "h":
		Help()
	case "default":
		// 没有参数，默认命令 Run
		Run()
	}
}
func Help() {
	println("帮助:\n可用参数(缩写):\n  构建: build(b)\n  运行: run(r)\n  压缩: UPX(u)\n  清理: clean(c)\n  帮助: help(h)\n例子:\n  make b\n  make build\n  make -b\n  make --build")
}
func Clean() {
	if isPathExists(pathExe) {
		err := os.Remove(pathExe)
		if err != nil {
			println("清理返回值：", err.Error())
		}
	}
}

func Build() bool {
	// 确认存在 go.exe
	pathGo, err := exec.LookPath("go")
	if err != nil {
		println("未找到 go 环境，请确认是否安装了 go 编译器")
		return false
	}
	println("构建...")
	// 设置环境变量
	cmd := exec.Command(pathGo, "build", "-ldflags", "-s -w", "-o", pathExe, pathSrc)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(
		os.Environ(),
		"GOOS=windows",
		"GOARCH=386",
	)
	println("命令:", cmd.String())
	println("构建输出:")
	err = cmd.Run()
	println("输出结束")
	if err != nil {
		println("构建返回值:", err.Error())
		return false
	} else {
		println("构建成功")
		return true
	}
}

func UPX() bool {
	if !isPathExists(pathUpx) {
		println("未找到UPX！跳过打包。")
		return false
	}
	println("UPX打包...")
	cmd := exec.Command(pathUpx, "-9", pathExe)
	err := cmd.Run()
	println("UPX打包结束")
	if err != nil {
		println("UPX返回值:", err.Error())
	}
	return true
}
func Run() bool {
	if !isPathExists(pathExe) {
		println("未找到EXE！开始构建")
		if Build() {
			println("=====运行输出=====")
		} else {
			return false
		}
	}
	cmd := exec.Command(pathExe)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	timeStart := time.Now()
	cmd.Run()
	println("执行用时:", time.Since(timeStart).String())
	return true
}

// 判断文件或文件夹是否存在,存在返回(true,nil),不存在返回(false,nil),未知返回(false,error)
func isPathExists(path string) bool {
	_, err := os.Stat(path)
	// 当文件或者文件夹存在
	if err == nil {
		return true
	}
	// 文件或文件夹不存在
	if os.IsNotExist(err) {
		return false
	}
	// 不确定
	return false
}
