package server

import (
	"chat/config"
	"log"
	http "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	mux "github.com/gorilla/mux"
)


type Server struct {
	cfg *config.Config
	router *mux.Router
}

func NewServer(cfg *config.Config) *Server {
	router := NewRouter(cfg.ApiVersion);
	return &Server{
		cfg: cfg,
		router: router,
	}
}


func (s *Server) Start() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
		WriteTimeout: time.Duration(s.cfg.SrvTimeout) * time.Second,
		ReadTimeout: time.Duration(s.cfg.SrvTimeout) * time.Second,
	}
	go func () {
		log.Println("Server is running on port 8080")
		log.Fatal(server.ListenAndServe())
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
}