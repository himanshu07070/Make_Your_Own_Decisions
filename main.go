package main

import (
	"Make_Your_Own_Decision/myod"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("file", "data.json", "this JSON file with MYOD story")
	port := flag.Int("port", 8080, "the port to start MYOD Web Application")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := myod.ParseJSON(f)
	if err != nil {
		panic(err)
	}

	h := myod.NewHandler(story, nil)
	fmt.Printf("Started server at Port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
