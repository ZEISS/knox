package main

import (
	"github.com/zeiss/knox/cmd"
)

func main() {
	err := cmd.Root.Execute()
	if err != nil {
		panic(err)
	}
}
