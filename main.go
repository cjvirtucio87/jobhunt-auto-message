package main

import (
	"./data"
	"log"
	"os"
	"text/template"
)

func main() {
	var cfg data.Config
	var f *os.File
	var err error
	cfg, err = data.FromJson("./config.json")

	tmpl := template.Must(template.New("message.txt").ParseFiles("message.txt"))

	f, err = os.OpenFile("build/result.txt", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = tmpl.Execute(f, cfg)

	if err != nil {
		log.Fatal(err)
	}
}
