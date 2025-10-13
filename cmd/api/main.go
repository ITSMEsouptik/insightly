package main

import (
	"github.com/ItsMeSouptik/insightly/internal/config"
	"github.com/ItsMeSouptik/insightly/internal/server"
)

func main() {
	cfg := config.Load()
	server.Run(cfg)
}
