package main

import (
	"fmt"

	"github.com/yatoenough/go-filecloud/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
