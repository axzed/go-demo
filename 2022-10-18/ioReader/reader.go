package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type AlphFilter struct {
	src string // 输入的字符串
	cur int    // 当前读取的位置
}

func alpha(r byte) byte {
	// r在 A-Z 或 a-z之间 不需要处理
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *AlphFilter) Read(p []byte) (int, error) {
	// 当前位置大于等于字符串的长度 说明读取到结尾了 返回EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}
	// 定义一个剩余还没有读到的长度
	x := len(a.src) - a.cur
	// bound 本次读取的长度
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else {
		bound = x
	}
	buf := make([]byte, bound)

	for n < bound {
		char := a.src[a.cur]
		res := alpha(a.src[a.cur])
		if res == 0 {
			continue
		}
		buf[n] = char
		n++
		a.cur++
	}
	copy(p, buf)
	return n, nil
}

func ReadFile() {
	fileName := "go.mod"
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("[内容: %s]", b)
}

func WriteFile() {
	fileName := "Daydream.txt"
	err := ioutil.WriteFile(fileName, []byte("我要成为技术大佬\n老婆孩子热炕头"), 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func ReadDir() {
	fs, err := ioutil.ReadDir("../")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range fs {
		fmt.Printf("[name:%v][size:%v][isDir:%v][mode:%v][ModTimeT%v]\n",
			f.Name(),	
			f.Size(),
			f.IsDir(),
			f.Mode(),
			f.ModTime(),
		)
	}
}

// 实现一个reader每次读取4字节
func main() {
	// 从字符串创建一个reader对象
	reader := strings.NewReader("vex web framework")
	// new 一个4字节的读取缓冲
	p := make([]byte, 4)
	for {
		// reader对象读数据
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				log.Printf("[数据已读完 EOF:%d]", n)
				break
			}
			log.Printf("[未知错误:%v]", err)
			return
		}
		log.Printf("[打印读取的字节数:%d 内容:%s]", n, string(p[:n]))
	}
	ReadFile()
	WriteFile()
	ReadDir()
}
