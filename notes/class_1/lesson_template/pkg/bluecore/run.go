package bluecore

import "fmt"

func HelloWorld(cfg Config) error {
	fmt.Printf("%s %s\n", cfg.Hello, cfg.World)
	return nil
}
