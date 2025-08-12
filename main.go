package main

import (
	"fmt"
	"github.com/Piep220/go-blog-aggregator/internal/config"
	"encoding/json"
)

func main() {
	fmt.Println("Starting...")
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	userName := currentUsername()
	cfg.SetUser(userName)

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	b, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(b))

}