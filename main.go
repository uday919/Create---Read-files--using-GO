package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Print("Welcome to files in Golang")
	content := "Hello, This is Uday and iam a Golang Developer. This is a sample file created by me and this program is about read and create files using Golang"
	file, err := os.Create("./sample.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is :", length)
	file.Close()
	read("./sample.txt")
}
func read(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data read from file is:", string(data))
}
