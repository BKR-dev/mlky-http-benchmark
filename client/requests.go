package client

import (
	"fmt"
	"net/http"
)

func main() {
	client()
}

func client() {
	resp, err := http.Get("")
	if err != nil {
		fmt.Println("ooope", err)
	}
	fmt.Println(resp)
}
