package main

import (
	"fmt"
	"log"
	"net/url"
	"no-go/registry"
)

func main() {
	u, err := url.Parse("https://registry.npmjs.org")
	if err != nil {
		log.Fatal(err)
	}

	r := registry.NewRegistry(u)

	m, err := r.Metadata("is-number")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.Name)
}
