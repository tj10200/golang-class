package main

import (
	"context"
	"fmt"
	"github.com/tj10200/golang-class/class_1/00_base_lesson_setup/cmd"
	"os"
)

func main() {
	root := cmd.NewRootCommand()
	if err := root.ExecuteContext(context.Background()); err != nil {
		fmt.Printf("root processing error: %v", err)
		os.Exit(1)
	}
}
