package main

import (
	"github.com/abdielrumaldo/aggro-gator/internal/config"
)

func main() {
	cfg := config.Config{}
	cfg.Read()
}
