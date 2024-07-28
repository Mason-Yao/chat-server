package main

import (
	"chat/config"
	"chat/internal/server"
)

// "context"
// "fmt"
// openai "github.com/sashabaranov/go-openai"
// "github.com/joho/godotenv"
// "os"

func main() {
	s := server.NewServer(config.NewConfig())
	s.Start();
}
