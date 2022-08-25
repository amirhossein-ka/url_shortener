package main

import (
	"log"

	"github.com/amirhossein-ka/url_shortener/cmd"
)

func main() {
    cmd.Parse()
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
}
