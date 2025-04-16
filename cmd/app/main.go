package main

import (
	"fmt"

	"github.com/liangtsao/demo-go-module-v2/internal/util"
	"github.com/liangtsao/demo-go-module-v2/pkg/calculator"
)

func main() {
	fmt.Println("Hello from demo-go-module-v2")

	reversed := util.ReverseString("Hello World")
	sum := calculator.Add(10, 20)
	fmt.Printf("Reversed: %s\nSum: %d\n", reversed, sum)
}
