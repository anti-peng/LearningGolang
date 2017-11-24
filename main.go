package main

import (
	"LearningGolang/toy7"
	"fmt"
)

func main() {
	str, err := toy7.DemoFillStd("ls", "-a")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}
