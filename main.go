package main

import (
	"flag"
	"os"
	"testing/app/http"
)
var (
	adrs = flag.String("adrs", os.Getenv("HTTP_ADRS"), "HTTP address to listen on")
)
func main() {
	flag.Parse()

	http.Serve(adrs)
}