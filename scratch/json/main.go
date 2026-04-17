package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	AllowedTools map[string]bool `json:"allowed_tools"`
}

func main() {
	data := []byte(`{"allowed_tools": {"hdn-server": true}}`)
	var cfg Config
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Config: %+v\n", cfg)
	for w, ok := range cfg.AllowedTools {
		fmt.Printf("w: %q, ok: %v\n", w, ok)
	}
}
