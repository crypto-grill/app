package main

import (
	"github.com/crypto-grill/app/internal/cli"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	if err := cli.Execute(os.Args); err != nil {
		_ = zap.S().Sync()
		log.Fatal(err)
	}
}
