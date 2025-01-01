package main

import (
	"fmt"
	"net/http"
)

func requestHandeller(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "testing world")
}
func main(){
	http.HandleFunc("/", requestHandeller)
	fmt.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
	}
}
