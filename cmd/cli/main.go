package main

import (
	"log"

	"github.com/ffelipelimao/mock-server-cli/internal/serialize"
	"github.com/ffelipelimao/mock-server-cli/internal/server"
	"github.com/ffelipelimao/mock-server-cli/internal/ui"
)

func main() {
	filePath, err := ui.GetFilePathUI()
	if err != nil {
		log.Fatal(err)
	}

	serialize := serialize.NewSerializer(filePath)

	config, err := serialize.Parse()
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(config)
	server.Start()
}
