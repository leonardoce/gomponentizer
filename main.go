// Package main is the entrypoint of the application
package main

import (
	"fmt"
	"os"

	"github.com/leonardoce/gomcoponentizer/cmd/gomcoponentizer"
)

func main() {
	err := gomcoponentizer.Cmd().Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
