package server

import (
	"log"
	"net/http"
	"os"

	"deepraj02/snoop/internal/handler"
	"deepraj02/snoop/pkg/network"
)

///`port`: The port number on which the server listens.
/// `dir`: The directory path where files are shared.
/// `handler`: A FileHandler instance to handle file requests.
type Server struct {
	port    string
	dir     string
	handler *handler.FileHandler
}

/// [Spawn] function creates a new Server object with the given port number and current directory path.
func Spawn(port string) *Server {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		port:    port,
		dir:     dir,
		handler: handler.New(dir),
	}
}

/// [Start] function starts the server on the specified port (8080 as default).
///
/// It registers the handler functions for the root path and the download path.

func (s *Server) Start() error {
	http.HandleFunc("/", s.handler.HandleIndex)
	http.HandleFunc("/download/", s.handler.HandleDownload)

	localIP, err := network.GetLocalNetworkIP()
	if err != nil {
		log.Printf("Warning: Could not determine local IP: %v", err)
		localIP = "localhost"
	}

	log.Printf("Starting server at http://%s:%s", localIP, s.port)
	return http.ListenAndServe(":"+s.port, nil)
}
