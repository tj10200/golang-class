package main

import (
	"context"
	"fmt"
	"github.com/tj10200/golang-class/notes/class_1/lesson_template/cmd"
	"os"
)

func main() {
	root := cmd.NewRootCommand()
	if err := root.ExecuteContext(context.Background()); err != nil {
		fmt.Printf("root processing error: %v", err)
		os.Exit(1)
	}
}
