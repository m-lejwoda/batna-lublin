package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Main Init")
	mux := http.NewServeMux()
	DatabaseURL := os.GetEnv("DATABASE_URL")	
}
