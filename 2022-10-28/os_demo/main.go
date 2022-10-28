package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i ++ {
		file.WriteString("WriteString写进来的\n")
		file.Write([]byte("Write写进来的\n"))
	}
}
