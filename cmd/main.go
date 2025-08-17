package main

import (
	"log"
	"os"

	"github.com/bugagych84/go1fl-sprint6-final-tpl/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse-converter ", log.LstdFlags|log.Lshortfile)
	s := server.New(logger)

	if err := s.Start(); err != nil {
		logger.Fatal(err)
	}
}
