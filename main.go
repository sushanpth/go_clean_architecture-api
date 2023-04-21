package main

import (
	"clean-architecture-api/bootstrap"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	_ = bootstrap.RootApp.Execute()
}
