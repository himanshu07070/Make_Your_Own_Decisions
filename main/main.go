package main

import (
	"Choose_Your_Own_Adventure/adventure"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("file", "choose_your.json", "this JSON file with choose_your_own story")
	port := flag.Int("port", 8080, "the port to start Choose_Your_Own_Story Web Application")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.ParseJSON(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story, nil)
	fmt.Printf("Started server at Port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
