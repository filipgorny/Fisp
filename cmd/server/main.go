package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8090", http.FileServer(http.Dir("./../../html/")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
