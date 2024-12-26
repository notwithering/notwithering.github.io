package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("."))); err != nil {
		fmt.Println(err)
		return
	}
}
