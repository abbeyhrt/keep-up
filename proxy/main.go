package main

import (
	"fmt"
	"os"
	"path"
)

type config struct{}

func main() {
	cfg := config{}

	// Paths
	gopath := os.Getenv("GOPATH")
	pkg := path.Join(gopath, "src/github.com/abbeyhrt/keep-up/proxy")
	src := path.Join(pkg, "nginx.conf.template")
	target := path.Join(pkg, "nginx.conf")

	// Template
	tmpl := template.New("nginx")
	tmpl, err := tmpl.ParseFiles(src)
	if err != nil {
		fmt.Errorf("error parsing files: %s", err)
	}

	f, err := os.Create(target)
	if err != nil {
		fmt.Errorf("error creating file: %s", err)
	}

	err = tmpl.Execute(f, nil)
	if err != nil {
		fmt.Errorf("error executing with config: %s", err)
	}
}
