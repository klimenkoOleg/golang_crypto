package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)

		written, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "io.Copy: reading %n: %v\n", written, err)
			os.Exit(1)
		}

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("Status code: %s\n", resp.Status)
		//fmt.Printf("%s", b)
	}
}
