package cmd

import (
	"github.com/tj10200/golang-class/class_1/00_base_lesson_setup/pkg/bluecore"
)

func RunCommand(cfg bluecore.Config) error {
	return bluecore.HelloWorld(cfg)
}
