package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.vercel.app/products"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close();
	// read 
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	fmt.Println(res.Status);


}