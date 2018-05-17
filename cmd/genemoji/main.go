package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/liuzl/dl"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	u = "https://raw.githubusercontent.com/matiassingers/emoji-flags/master/data.json"
	d = "src/github.com/liuzl/gocountries/emojis.go"
)

type Item struct {
	Code  string `json:"code"`
	Emoji string `json:"emoji"`
}

func main() {
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("Missing $GOPATH environment variable")
	}
	path := filepath.Join(gopath, d)

	resp := dl.DownloadUrl(u)
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	var items []Item
	if err := json.Unmarshal(resp.Content, &items); err != nil {
		log.Fatal(err)
	}
	output := bytes.Buffer{}
	output.WriteString("package gocountries\n\n")
	output.WriteString("var Emojis = map[string]string{\n")
	for _, item := range items {
		output.WriteString(fmt.Sprintf("\t\"%s\": \"%s\",\n", item.Code, item.Emoji))
	}
	output.WriteString("}\n")
	if err := ioutil.WriteFile(path, output.Bytes(), os.FileMode(0644)); err != nil {
		log.Fatal(err)
	}
}
