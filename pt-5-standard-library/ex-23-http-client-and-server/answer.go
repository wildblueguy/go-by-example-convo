package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Configure default router
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(responseWriter, "Hi")
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Println(err)
		}
	}()
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
