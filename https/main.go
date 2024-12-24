package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://jsonplaceholder.typicode.com/posts/1");
	if err !=nil {
		log.Fatalf("Error is found", err);
	}
	defer resp.Body.Close();

	body , err := ioutil.ReadAll(resp.Body);
	if err !=nil{
		log.Fatalf("Error is found", err);
	}

	fmt.Println("res stataus", resp.Status)
	fmt.Println("res body", string(body))

}