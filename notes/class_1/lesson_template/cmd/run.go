package cmd

import (
	"github.com/tj10200/golang-class/notes/class_1/lesson_template/pkg/bluecore"
)

func RunCommand(cfg bluecore.Config) error {
	return bluecore.HelloWorld(cfg)
}
