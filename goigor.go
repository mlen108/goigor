package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const listUrl = "https://api.github.com/repos/github/gitignore/contents"

type Tmpl struct {
	Name string `json:"name"`
	Url  string `json:"download_url"`
}

func main() {
	helpUsage := `
    Usage: goigor [name...]

    Available names: `

	templates, err := list()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		var avblNames []string
		for _, t := range templates {
			avblNames = append(avblNames, t.Name)
		}

		fmt.Printf("%s%s.", helpUsage, strings.Join(avblNames, ", "))
		os.Exit(0)
	}

	names := os.Args[1:]
	for _, n := range names {
		for _, t := range templates {
			if n == t.Name {
				content, err := fetch(t.Url)
				if err != nil {
					fmt.Printf("Error: %s", err)
					return
				}
				b := bytes.NewBuffer(content)
				b.WriteTo(os.Stdout)
				continue
			}
		}
	}
}

func fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request failed: HTTP%d", res.StatusCode))
	}

	return ioutil.ReadAll(res.Body)
}

func list() ([]Tmpl, error) {
	var templates []Tmpl
	files := []Tmpl{}

	res, err := fetch(listUrl)
	if err != nil {
		return nil, err
	}
	resBuff := bytes.NewBuffer(res)
	json.NewDecoder(resBuff).Decode(&templates)

	for _, t := range templates {
		if strings.HasSuffix(t.Name, ".gitignore") {
			name := strings.ToLower(t.Name[:len(t.Name)-10])
			files = append(files, Tmpl{name, t.Url})
		}
	}
	return files, nil
}
