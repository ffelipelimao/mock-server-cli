package server

import (
	"fmt"

	"github.com/ffelipelimao/mock-server-cli/internal/serialize"
)

type Server struct {
	config *serialize.Config
}

func NewServer(config *serialize.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() {
	fmt.Println("Sever Running...")
	for {
	}
}
