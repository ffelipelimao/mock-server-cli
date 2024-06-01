package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ffelipelimao/mock-server-cli/internal/serialize"
	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()

	for _, req := range s.config.Requests {
		req := req // avoid closure capturing loop variable

		switch strings.ToUpper(req.Verb) {
		case "GET":
			fmt.Printf("Router - Verb: %s Path %s \n", req.Verb, req.Path)
			router.Get(req.Path, s.createHandler(req))
		case "POST":
			fmt.Printf("Router - Verb: %s Path %s \n", req.Verb, req.Path)
			router.Post(req.Path, s.createHandler(req))
		case "PUT":
			fmt.Printf("Router - Verb: %s Path %s \n", req.Verb, req.Path)
			router.Put(req.Path, s.createHandler(req))
		case "PATCH":
			fmt.Printf("Router - Verb: %s Path %s \n", req.Verb, req.Path)
			router.Patch(req.Path, s.createHandler(req))
		}
	}

	serverPort := fmt.Sprintf(":%s", s.config.Port)

	fmt.Println("Server Running...")
	http.ListenAndServe(serverPort, router)
}

func (s *Server) createHandler(req serialize.Requests) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := s.prepareResponse(req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		contentType := fmt.Sprintf("application/%s", strings.ToLower(req.ResponseType))
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func (s *Server) prepareResponse(req serialize.Requests) ([]byte, error) {
	if strings.ToLower(req.ResponseType) == "xml" {
		if str, ok := req.Response.Body.(string); ok {
			return []byte(str), nil
		} else {
			return nil, fmt.Errorf("invalid XML response body")
		}
	}
	return json.Marshal(req.Response.Body)
}
