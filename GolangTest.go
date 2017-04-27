package main

import (
	"fmt"
	"os"
)

func main() {
	err1 := os.Chdir("F:/")
	if err1 != nil {
		fmt.Println(err1)
	}
	dir, err2 := os.Getwd()
	if err2 == nil {
		fmt.Println("当前目录：", dir)
	}

}
