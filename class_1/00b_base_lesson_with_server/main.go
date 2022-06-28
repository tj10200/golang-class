package main

import (
	"context"
	"fmt"
	base "github.com/tj10200/golang-class/class_1/00_base_lesson_setup/cmd"
	"github.com/tj10200/golang-class/class_1/00b_base_lesson_with_server/cmd"
	"os"
)

func main() {
	root := base.NewRootCommand()
	root.AddCommand(cmd.NewServerCommand())
	if err := root.ExecuteContext(context.Background()); err != nil {
		fmt.Printf("root processing error: %v", err)
		os.Exit(1)
	}
}
