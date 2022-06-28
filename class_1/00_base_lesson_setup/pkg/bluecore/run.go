package bluecore

import (
	"encoding/json"
	"fmt"
)

func HelloWorld(cfg Config) error {
	fmt.Printf("%s %s %+v\n", cfg.Hello, cfg.World, cfg)
	data, err := json.Marshal(&cfg)
	if err != nil {
		return fmt.Errorf("ran into trouble marshalling json: %v", err)
	}
	fmt.Printf("config received: %s\n", string(data))
	return nil
}
