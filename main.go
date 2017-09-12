package main

import (
	"github.com/omise/omise-go"
	"net/http"
	"os"
)

func main() {
	key := os.Getenv("OMISE_SKEY")
	if key == "" {
		panic("Please set OMISE_SKEY")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("Please set PORT")
	}

	client, e := omise.NewClient("", key)
	if e != nil {
		panic(e)
	}

	if e := http.ListenAndServe(":" + port, &TamboonHandler{client}); e != nil {
		panic(e)
	}
}
