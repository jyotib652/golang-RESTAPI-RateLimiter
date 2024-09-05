package main

import (
	"fmt"
	"net/http"
)

const webPort = 8080

type JSONResponse struct {
	Error   bool
	Message string
}

func main() {
	fmt.Printf("Web server is starting on http://localhost:%d\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
