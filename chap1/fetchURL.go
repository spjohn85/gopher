package main

//Fetches URL using GET method
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(resp.Status)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch url: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}
