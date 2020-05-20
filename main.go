package main

import (
	"./AddGoshed"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// src is the input for which we want to print the AST.
	var fileName string
	fmt.Fscan(os.Stdin, &fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print(fmt.Errorf("unable to open file %v: %v", fileName, err))
		panic(err)
	}
	defer file.Close()
	src, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(fmt.Errorf("unable to read from file %v: %v", fileName, err))
		panic(err)
	}
	newProgram := AddGoshed.AddGoschedToFile("", string(src))
	err = ioutil.WriteFile("output.go", []byte(newProgram), 0777)
	if err != nil {
		fmt.Println(err)
	}
}