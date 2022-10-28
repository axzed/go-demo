package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("[获取命令行参数][res:%v]", os.Args)
	name, _ := os.Hostname()
	log.Printf("[获取主机名][res:%v]", name)
	log.Printf("[获取当前进程名][res:%v]", os.Getpid())
	log.Printf("[获取一条环境变量][res:%v]", os.Getenv("GOROOT"))

	// 获取所有环境变量
	env := os.Environ()
	for _, v := range env {
		log.Printf("[获取所有环境变量][res:%v]", v)
	}
	dir, _ := os.Getwd()
	log.Printf("[获取当前目录][res:%v]", dir)
}
